#!/bin/bash
set -euo pipefail

USAGE="Runs the action runner against a local git repo.

This is mainly useful for testing the action runner locally without needing
to manually craft the exact webhook payloads that would be needed to trigger
it correctly.

Examples:
  # Run actions for the BuildBuddy repo
  $0

  # Run actions for a local repo
  REPO_PATH=~/src/scratch/hello_world $0

  # Override CI runner args
  $0 --bes_backend=grpcs://cloud.buildbuddy.io --bes_results_url=https://app.buildbuddy.io/invocation/
"
if [[ "${1:-}" =~ ^(-h|--help)$ ]]; then
  echo "$USAGE"
  exit
fi

# cd to workspace root
cd "$(dirname "$0")"
while ! [ -e "WORKSPACE" ]; do
  cd ..
  if [[ "$PWD" == / ]]; then
    echo >&2 "Failed to find the bazel workspace root containing this script."
    exit 1
  fi
done

dir_abspath() (cd "$1" && pwd)

# CI runner bazel cache is set to a fixed directory in order
# to speed up builds, but note that in production we don't yet
# have persistent local caching.
TEMPDIR=$(mktemp --dry-run | xargs dirname)
: "${CI_RUNNER_BAZEL_CACHE_DIR:=$TEMPDIR/buildbuddy_ci_runner_bazel_cache}"

: "${REPO_PATH:=$PWD}"

bazel build //enterprise/server/cmd/ci_runner:buildbuddy_ci_runner
RUNNER_PATH="$PWD/bazel-bin/enterprise/server/cmd/ci_runner/buildbuddy_ci_runner"
echo "$RUNNER_PATH"

mkdir -p "$CI_RUNNER_BAZEL_CACHE_DIR"

docker run \
  --volume "$RUNNER_PATH:/bin/ci_runner" \
  --volume "$CI_RUNNER_BAZEL_CACHE_DIR:/root/.cache/bazel" \
  --volume "$(dir_abspath "$REPO_PATH"):/root/mounted_repo" \
  --net host \
  --interactive \
  --tty \
  --rm \
  gcr.io/flame-public/buildbuddy-ci-runner:v2.2.7 \
  ci_runner \
  --pushed_repo_url="file:///root/mounted_repo" \
  --target_repo_url="file:///root/mounted_repo" \
  --commit_sha="$(cd "$REPO_PATH" && git rev-parse HEAD)" \
  --pushed_branch="$(cd "$REPO_PATH" && git branch --show-current)" \
  --target_branch="master" \
  --trigger_event=pull_request \
  --bes_backend=grpc://localhost:1985 \
  --bes_results_url=http://localhost:8080/invocation/ \
  "$@"
