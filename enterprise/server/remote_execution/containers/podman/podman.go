package podman

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/buildbuddy-io/buildbuddy/enterprise/server/remote_execution/commandutil"
	"github.com/buildbuddy-io/buildbuddy/enterprise/server/remote_execution/container"
	"github.com/buildbuddy-io/buildbuddy/server/environment"
	"github.com/buildbuddy-io/buildbuddy/server/interfaces"
	"github.com/buildbuddy-io/buildbuddy/server/util/alert"
	"github.com/buildbuddy-io/buildbuddy/server/util/background"
	"github.com/buildbuddy-io/buildbuddy/server/util/grpc_client"
	"github.com/buildbuddy-io/buildbuddy/server/util/log"
	"github.com/buildbuddy-io/buildbuddy/server/util/lru"
	"github.com/buildbuddy-io/buildbuddy/server/util/networking"
	"github.com/buildbuddy-io/buildbuddy/server/util/perms"
	"github.com/buildbuddy-io/buildbuddy/server/util/random"
	"github.com/buildbuddy-io/buildbuddy/server/util/status"

	regpb "github.com/buildbuddy-io/buildbuddy/proto/registry"
	repb "github.com/buildbuddy-io/buildbuddy/proto/remote_execution"
)

var (
	imageStreamingRegistryGRPCTarget = flag.String("executor.podman_image_streaming.registry_grpc_target", "", "gRPC endpoint of BuildBuddy registry")
	imageStreamingRegistryHTTPTarget = flag.String("executor.podman_image_streaming.registry_http_target", "", "HTTP endpoint of the BuildBuddy registry")

	// Additional time used to kill the container if the command doesn't exit cleanly
	containerFinalizationTimeout = 10 * time.Second

	storageErrorRegex = regexp.MustCompile(`(?s)A storage corruption might have occurred.*Error: readlink.*no such file or directory`)
	userRegex         = regexp.MustCompile(`^[a-z0-9][a-z0-9_-]*(:[a-z0-9_-]*)?$`)

	// A map from image name to pull status. This is used to avoid parallel pulling of the same image.
	pullOperations sync.Map
)

const (
	podmanInternalExitCode = 125
	// podmanExecSIGKILLExitCode is the exit code returned by `podman exec` when the exec
	// process is killed due to the parent container being removed.
	podmanExecSIGKILLExitCode = 137

	podmanDefaultNetworkIPRange = "10.88.0.0/16"
	podmanDefaultNetworkGateway = "10.88.0.1"
	podmanDefaultNetworkBridge  = "cni-podman0"

	optImageRefCacheSize = 1000
)

type optImageCache struct {
	cache interfaces.LRU
}

func (c *optImageCache) get(key string) (string, error) {
	if v, ok := c.cache.Get(key); ok {
		return v.(string), nil
	}
	return "", nil
}

func (c *optImageCache) put(key string, optImage string) {
	c.cache.Add(key, optImage)
}

func newOptImageCache() (*optImageCache, error) {
	l, err := lru.NewLRU(&lru.Config{
		SizeFn: func(value interface{}) int64 {
			return 1
		},
		MaxSize: optImageRefCacheSize,
	})
	if err != nil {
		return nil, err
	}
	return &optImageCache{cache: l}, nil
}

type pullStatus struct {
	mu     *sync.RWMutex
	pulled bool
}

type Provider struct {
	env                   environment.Env
	imageCacheAuth        *container.ImageCacheAuthenticator
	buildRoot             string
	imageStreamingEnabled bool
	regClient             regpb.RegistryClient
	optImageCache         *optImageCache
}

func NewProvider(env environment.Env, imageCacheAuthenticator *container.ImageCacheAuthenticator, buildRoot string) (*Provider, error) {
	c, err := newOptImageCache()
	if err != nil {
		return nil, err
	}
	var regClient regpb.RegistryClient
	if *imageStreamingRegistryGRPCTarget != "" {
		conn, err := grpc_client.DialTarget(*imageStreamingRegistryGRPCTarget)
		if err != nil {
			return nil, err
		}
		regClient = regpb.NewRegistryClient(conn)
	}

	imageStreamingEnabled := *imageStreamingRegistryHTTPTarget != "" && *imageStreamingRegistryGRPCTarget != ""
	if imageStreamingEnabled {
		log.Infof("Starting stargz store")
		cmd := exec.CommandContext(env.GetServerContext(), "stargz-store", "/var/lib/stargz-store/store")
		logWriter := log.Writer("[stargzstore] ")
		cmd.Stderr = logWriter
		cmd.Stdout = logWriter
		if err := cmd.Start(); err != nil {
			return nil, status.UnavailableErrorf("could not start stargz store: %s", err)
		}

		// Configures podman to check stargz store for image data.
		storageConf := `
[storage]
driver = "overlay"
runroot = "/run/containers/storage"
graphroot = "/var/lib/containers/storage"
[storage.options]
additionallayerstores=["/var/lib/stargz-store/store:ref"]
`
		if err := os.WriteFile("/etc/containers/storage.conf", []byte(storageConf), 0644); err != nil {
			return nil, status.UnavailableErrorf("could not write storage config: %s", err)
		}
	}

	return &Provider{
		env:                   env,
		imageCacheAuth:        imageCacheAuthenticator,
		buildRoot:             buildRoot,
		imageStreamingEnabled: imageStreamingEnabled,
		regClient:             regClient,
		optImageCache:         c,
	}, nil
}

func (p *Provider) NewContainer(image string, options *PodmanOptions) container.CommandContainer {
	return &podmanCommandContainer{
		env:                   p.env,
		imageCacheAuth:        p.imageCacheAuth,
		image:                 image,
		registryClient:        p.regClient,
		imageStreamingEnabled: *imageStreamingRegistryGRPCTarget != "" && *imageStreamingRegistryHTTPTarget != "",
		optImageCache:         p.optImageCache,
		buildRoot:             p.buildRoot,
		options:               options,
	}
}

type PodmanOptions struct {
	ForceRoot bool
	User      string
	Network   string
	CapAdd    string
	Devices   []container.DockerDeviceMapping
	Volumes   []string
	Runtime   string
}

// podmanCommandContainer containerizes a command's execution using a Podman container.
// between containers.
type podmanCommandContainer struct {
	env            environment.Env
	imageCacheAuth *container.ImageCacheAuthenticator

	image     string
	buildRoot string

	imageStreamingEnabled bool
	optImageCache         *optImageCache
	registryClient        regpb.RegistryClient
	// If container streaming is enabled, this field will contain the name of
	// the image optimized for streaming.
	optimizedImage string

	options *PodmanOptions

	// name is the container name.
	name string

	mu sync.Mutex // protects(removed)
	// removed is a flag that is set once Remove is called (before actually
	// removing the container).
	removed bool
}

func NewPodmanCommandContainer(env environment.Env, imageCacheAuth *container.ImageCacheAuthenticator, image, buildRoot string, options *PodmanOptions) container.CommandContainer {
	return &podmanCommandContainer{
		env:            env,
		imageCacheAuth: imageCacheAuth,
		image:          image,
		buildRoot:      buildRoot,
		options:        options,
	}
}

func (c *podmanCommandContainer) getPodmanRunArgs(workDir string) []string {
	args := []string{
		"--hostname",
		"localhost",
		"--workdir",
		workDir,
		"--name",
		c.name,
		"--rm",
		"--dns",
		"8.8.8.8",
		"--dns-search",
		".",
		"--volume",
		fmt.Sprintf(
			"%s:%s",
			filepath.Join(c.buildRoot, filepath.Base(workDir)),
			workDir,
		),
	}
	if c.options.ForceRoot {
		args = append(args, "--user=0:0")
	} else if c.options.User != "" && userRegex.MatchString(c.options.User) {
		args = append(args, "--user="+c.options.User)
	}
	if strings.ToLower(c.options.Network) == "off" {
		args = append(args, "--network=none")
	}
	if c.options.CapAdd != "" {
		args = append(args, "--cap-add="+c.options.CapAdd)
	}
	for _, device := range c.options.Devices {
		deviceSpecs := make([]string, 0)
		if device.PathOnHost != "" {
			deviceSpecs = append(deviceSpecs, device.PathOnHost)
		}
		if device.PathInContainer != "" {
			deviceSpecs = append(deviceSpecs, device.PathInContainer)
		}
		if device.CgroupPermissions != "" {
			deviceSpecs = append(deviceSpecs, device.CgroupPermissions)
		}
		args = append(args, "--device="+strings.Join(deviceSpecs, ":"))
	}
	for _, volume := range c.options.Volumes {
		args = append(args, "--volume="+volume)
	}
	if c.options.Runtime != "" {
		args = append(args, "--runtime="+c.options.Runtime)
	}
	return args
}

func (c *podmanCommandContainer) Run(ctx context.Context, command *repb.Command, workDir string, creds container.PullCredentials) *interfaces.CommandResult {
	result := &interfaces.CommandResult{
		CommandDebugString: fmt.Sprintf("(podman) %s", command.GetArguments()),
		ExitCode:           commandutil.NoExitCode,
	}
	containerName, err := generateContainerName()
	c.name = containerName
	if err != nil {
		result.Error = status.UnavailableErrorf("failed to generate podman container name: %s", err)
		return result
	}
	if err := container.PullImageIfNecessary(ctx, c.env, c.imageCacheAuth, c, creds, c.image); err != nil {
		result.Error = status.UnavailableErrorf("failed to pull docker image: %s", err)
		return result
	}

	image, err := c.targetImage(ctx)
	if err != nil {
		result.Error = err
		return result
	}

	podmanRunArgs := c.getPodmanRunArgs(workDir)

	for _, envVar := range command.GetEnvironmentVariables() {
		podmanRunArgs = append(podmanRunArgs, "--env", fmt.Sprintf("%s=%s", envVar.GetName(), envVar.GetValue()))
	}
	podmanRunArgs = append(podmanRunArgs, image)
	podmanRunArgs = append(podmanRunArgs, command.Arguments...)
	result = runPodman(ctx, "run", &container.ExecOpts{}, podmanRunArgs...)
	if err := c.maybeCleanupCorruptedImages(ctx, result); err != nil {
		log.Warningf("Failed to remove corrupted image: %s", err)
	}
	if exitedCleanly := result.ExitCode >= 0; !exitedCleanly {
		if err = c.killContainerIfRunning(ctx); err != nil {
			log.Warningf("Failed to shut down podman container: %s", err)
		}
	}
	return result
}

func (c *podmanCommandContainer) optImageRefKey(ctx context.Context) (string, error) {
	groupID := ""
	u, err := perms.AuthenticatedUser(ctx, c.env)
	if err != nil {
		if !perms.IsAnonymousUserError(err) {
			return "", err
		}
	} else {
		groupID = u.GetGroupID()
	}

	return fmt.Sprintf("%s-%s", groupID, c.image), nil
}

func (c *podmanCommandContainer) targetImage(ctx context.Context) (string, error) {
	if !c.imageStreamingEnabled {
		return c.image, nil
	}

	// If the optimized image name has not been resolved from the registry
	// yet then check our local cache.
	if c.optimizedImage == "" {
		key, err := c.optImageRefKey(ctx)
		if err != nil {
			return "", err
		}
		optImage, err := c.optImageCache.get(key)
		if err != nil {
			return "", err
		}
		c.optimizedImage = optImage
	}
	if c.optimizedImage == "" {
		return "", status.FailedPreconditionErrorf("optimized image not yet resolved")
	}
	return c.optimizedImage, nil
}

func (c *podmanCommandContainer) resolveTargetImage(ctx context.Context, credentials container.PullCredentials) (string, error) {
	if !c.imageStreamingEnabled {
		return c.image, nil
	}

	if c.registryClient == nil || *imageStreamingRegistryHTTPTarget == "" {
		return "", status.FailedPreconditionErrorf("streaming enabled, but registry client or http target are not set")
	}

	log.CtxInfof(ctx, "Resolving optimized image for %q", c.image)
	req := &regpb.GetOptimizedImageRequest{
		Image: c.image,
		Platform: &regpb.Platform{
			Arch: runtime.GOARCH,
			Os:   runtime.GOOS,
		},
	}
	if !credentials.IsEmpty() {
		req.ImageCredentials = &regpb.Credentials{
			Username: credentials.Username,
			Password: credentials.Password,
		}
	}

	// Clear the JWT from the RPC so that the blobs are stored as anonymous data
	// until we implement CAS auth for image streaming.
	rpcCtx := context.WithValue(ctx, "x-buildbuddy-jwt", nil)
	rsp, err := c.registryClient.GetOptimizedImage(rpcCtx, req)
	if err != nil {
		return "", status.UnavailableErrorf("could not resolve optimized image for %q: %s", c.image, err)
	}
	optImage := fmt.Sprintf("%s/%s", *imageStreamingRegistryHTTPTarget, rsp.GetOptimizedImage())
	log.CtxInfof(ctx, "Resolved optimized image %q for %q", optImage, c.image)
	c.optimizedImage = optImage
	key, err := c.optImageRefKey(ctx)
	c.optImageCache.put(key, optImage)
	return c.optimizedImage, nil
}

func (c *podmanCommandContainer) Create(ctx context.Context, workDir string) error {
	containerName, err := generateContainerName()
	if err != nil {
		return status.UnavailableErrorf("failed to generate podman container name: %s", err)
	}
	c.name = containerName

	podmanRunArgs := c.getPodmanRunArgs(workDir)
	image, err := c.targetImage(ctx)
	if err != nil {
		return err
	}
	podmanRunArgs = append(podmanRunArgs, image)
	podmanRunArgs = append(podmanRunArgs, "sleep", "infinity")
	createResult := runPodman(ctx, "create", &container.ExecOpts{}, podmanRunArgs...)
	if err := c.maybeCleanupCorruptedImages(ctx, createResult); err != nil {
		log.Warningf("Failed to remove corrupted image: %s", err)
	}

	if err = createResult.Error; err != nil {
		return status.UnavailableErrorf("failed to create container: %s", err)
	}

	if createResult.ExitCode != 0 {
		return status.UnknownErrorf("podman create failed: exit code %d, stderr: %s", createResult.ExitCode, createResult.Stderr)
	}

	startResult := runPodman(ctx, "start", &container.ExecOpts{}, c.name)
	if startResult.Error != nil {
		return startResult.Error
	}
	if startResult.ExitCode != 0 {
		return status.UnknownErrorf("podman start failed: exit code %d, stderr: %s", startResult.ExitCode, startResult.Stderr)
	}
	return nil
}

func (c *podmanCommandContainer) Exec(ctx context.Context, cmd *repb.Command, opts *container.ExecOpts) *interfaces.CommandResult {
	podmanRunArgs := make([]string, 0, 2*len(cmd.GetEnvironmentVariables())+len(cmd.Arguments)+1)
	for _, envVar := range cmd.GetEnvironmentVariables() {
		podmanRunArgs = append(podmanRunArgs, "--env", fmt.Sprintf("%s=%s", envVar.GetName(), envVar.GetValue()))
	}
	if c.options.ForceRoot {
		podmanRunArgs = append(podmanRunArgs, "--user=0:0")
	} else if c.options.User != "" && userRegex.MatchString(c.options.User) {
		podmanRunArgs = append(podmanRunArgs, "--user="+c.options.User)
	}
	if strings.ToLower(c.options.Network) == "off" {
		podmanRunArgs = append(podmanRunArgs, "--network=none")
	}
	if opts.Stdin != nil {
		podmanRunArgs = append(podmanRunArgs, "--interactive")
	}
	podmanRunArgs = append(podmanRunArgs, c.name)
	podmanRunArgs = append(podmanRunArgs, cmd.Arguments...)
	// Podman doesn't provide a way to find out whether an exec process was
	// killed. Instead, `podman exec` returns 137 (= 128 + SIGKILL(9)). However,
	// this exit code is also valid as a regular exit code returned by a command
	// during a normal execution, so we are overly cautious here and only
	// interpret this code specially when the container was removed and we are
	// expecting a SIGKILL as a result.
	res := runPodman(ctx, "exec", opts, podmanRunArgs...)
	c.mu.Lock()
	removed := c.removed
	c.mu.Unlock()
	if removed && res.ExitCode == podmanExecSIGKILLExitCode {
		res.ExitCode = commandutil.KilledExitCode
		res.Error = commandutil.ErrSIGKILL
	}
	return res
}

func (c *podmanCommandContainer) IsImageCached(ctx context.Context) (bool, error) {
	if c.imageStreamingEnabled {
		return true, nil
	}

	// Try to avoid the `pull` command which results in a network roundtrip.
	listResult := runPodman(ctx, "image", &container.ExecOpts{}, "inspect", "--format={{.ID}}", c.image)
	if listResult.ExitCode == podmanInternalExitCode {
		return false, nil
	} else if listResult.Error != nil {
		return false, listResult.Error
	}

	if strings.TrimSpace(string(listResult.Stdout)) != "" {
		// Found at least one image matching the ref; `docker run` should succeed
		// without pulling the image.
		return true, nil
	}
	return false, nil
}

func (c *podmanCommandContainer) PullImage(ctx context.Context, creds container.PullCredentials) error {
	psi, _ := pullOperations.LoadOrStore(c.image, &pullStatus{&sync.RWMutex{}, false})
	ps, ok := psi.(*pullStatus)
	if !ok {
		alert.UnexpectedEvent("psi cannot be cast to *pullStatus")
		return status.InternalError("PullImage failed: cannot get pull status")
	}

	ps.mu.RLock()
	alreadyPulled := ps.pulled
	ps.mu.RUnlock()

	if alreadyPulled {
		return c.pullImage(ctx, creds)
	}

	ps.mu.Lock()
	defer ps.mu.Unlock()
	if err := c.pullImage(ctx, creds); err != nil {
		return err
	}
	ps.pulled = true
	return nil
}

func (c *podmanCommandContainer) pullImage(ctx context.Context, creds container.PullCredentials) error {
	if c.imageStreamingEnabled {
		// Always re-resolve image when a pull is requested. This takes care of
		// re-validating the passed credentials.
		if _, err := c.resolveTargetImage(ctx, creds); err != nil {
			return err
		}
		return nil
	}

	podmanArgs := make([]string, 0, 2)
	if !creds.IsEmpty() {
		podmanArgs = append(podmanArgs, fmt.Sprintf(
			"--creds=%s:%s",
			creds.Username,
			creds.Password,
		))
	}
	podmanArgs = append(podmanArgs, c.image)
	// Use server context instead of ctx to make sure that "podman pull" is not killed when the context
	// is cancelled. If "podman pull" is killed when copying a parent layer, it will result in
	// corrupted storage.  More details see https://github.com/containers/storage/issues/1136.
	pullResult := runPodman(c.env.GetServerContext(), "pull", &container.ExecOpts{}, podmanArgs...)
	if pullResult.Error != nil {
		return pullResult.Error
	}
	if pullResult.ExitCode != 0 {
		return status.UnknownErrorf("podman pull failed: exit code %d, stderr: %s", pullResult.ExitCode, string(pullResult.Stderr))
	}
	return nil
}

func (c *podmanCommandContainer) Remove(ctx context.Context) error {
	c.mu.Lock()
	c.removed = true
	c.mu.Unlock()
	res := runPodman(ctx, "kill", &container.ExecOpts{}, "--signal=KILL", c.name)
	if res.Error != nil {
		return res.Error
	}
	if res.ExitCode == 0 || strings.Contains(string(res.Stderr), "no such container") {
		return nil
	}
	return status.UnknownErrorf("podman remove failed: exit code %d, stderr: %s", res.ExitCode, string(res.Stderr))
}

func (c *podmanCommandContainer) Pause(ctx context.Context) error {
	res := runPodman(ctx, "pause", &container.ExecOpts{}, c.name)
	if res.ExitCode != 0 {
		return status.UnknownErrorf("podman pause failed: exit code %d, stderr: %s", res.ExitCode, string(res.Stderr))
	}
	return nil
}

func (c *podmanCommandContainer) Unpause(ctx context.Context) error {
	res := runPodman(ctx, "unpause", &container.ExecOpts{}, c.name)
	if res.Error != nil {
		return res.Error
	}
	if res.ExitCode != 0 {
		return status.UnknownErrorf("podman unpause failed: exit code %d, stderr: %s", res.ExitCode, string(res.Stderr))
	}
	return nil
}

func (c *podmanCommandContainer) Stats(ctx context.Context) (*container.Stats, error) {
	return &container.Stats{}, nil
}

func runPodman(ctx context.Context, subCommand string, opts *container.ExecOpts, args ...string) *interfaces.CommandResult {
	command := []string{
		"podman",
		subCommand,
	}

	command = append(command, args...)
	result := commandutil.Run(ctx, &repb.Command{Arguments: command}, "" /*=workDir*/, opts)
	return result
}

func generateContainerName() (string, error) {
	suffix, err := random.RandomString(20)
	if err != nil {
		return "", err
	}
	return "buildbuddy_exec_" + suffix, nil
}

func (c *podmanCommandContainer) killContainerIfRunning(ctx context.Context) error {
	ctx, cancel := background.ExtendContextForFinalization(ctx, containerFinalizationTimeout)
	defer cancel()

	err := c.Remove(ctx)
	if err != nil && strings.Contains(err.Error(), "Error: can only kill running containers.") {
		// This is expected.
		return nil
	}
	return err
}

// An image can be corrupted if "podman pull" command is killed when pulling a parent layer.
// More details can be found at https://github.com/containers/storage/issues/1136. When this
// happens when need to remove the image before re-pulling the image in order to fix it.
func (c *podmanCommandContainer) maybeCleanupCorruptedImages(ctx context.Context, result *interfaces.CommandResult) error {
	if result.ExitCode != podmanInternalExitCode {
		return nil
	}
	if !storageErrorRegex.MatchString(string(result.Stderr)) {
		return nil
	}
	result.Error = status.UnavailableError("a storage corruption occurred")
	result.ExitCode = commandutil.NoExitCode
	return removeImage(ctx, c.image)
}

func removeImage(ctx context.Context, imageName string) error {
	ctx, cancel := background.ExtendContextForFinalization(ctx, containerFinalizationTimeout)
	defer cancel()

	result := runPodman(ctx, "rmi", &container.ExecOpts{}, imageName)
	if result.Error != nil {
		return result.Error
	}
	if result.ExitCode == 0 || strings.Contains(string(result.Stderr), "image not known") {
		return nil
	}
	return status.UnknownErrorf("podman rmi failed: %s", string(result.Stderr))
}

// Configure the secondary network for podman so that traffic from podman will be routed through
// the secondary network interface instead of the primary network.
func ConfigureSecondaryNetwork(ctx context.Context) error {
	if !networking.IsSecondaryNetworkEnabled() {
		// No need to configure secondary network for podman.
		return nil
	}
	// Hack: run a dummy podman container to setup default podman bridge network in ip route.
	// "podman run --rm busybox sh". This should setup the following in ip route:
	// "10.88.0.0/16 dev cni-podman0 proto kernel scope link src 10.88.0.1 linkdown"
	result := runPodman(ctx, "run", &container.ExecOpts{}, "--rm", "busybox", "sh")
	if result.Error != nil {
		return result.Error
	}
	if result.ExitCode != 0 {
		return status.UnknownError("failed to setup podman default network")
	}

	// Add ip rule to lookup rt1
	// Equivalent to "ip rule add to 10.88.0.0/16 lookup rt1"
	if err := networking.AddIPRuleIfNotPresent(ctx, []string{"to", podmanDefaultNetworkIPRange}); err != nil {
		return err
	}
	if err := networking.AddIPRuleIfNotPresent(ctx, []string{"from", podmanDefaultNetworkIPRange}); err != nil {
		return err
	}

	// Add ip route to routing table rt1
	// Equivalent to "ip route add 10.88.0.0/16 via 10.88.0.1 dev cni-podman0 table rt1"
	route := []string{podmanDefaultNetworkIPRange, "via", podmanDefaultNetworkGateway, "dev", podmanDefaultNetworkBridge}
	if err := networking.AddRouteIfNotPresent(ctx, route); err != nil {
		return err
	}
	return nil

}
