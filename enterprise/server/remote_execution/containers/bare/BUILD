load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "bare",
    srcs = ["bare.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/remote_execution/containers/bare",
    visibility = ["//visibility:public"],
    deps = [
        "//enterprise/server/remote_execution/commandutil",
        "//enterprise/server/remote_execution/container",
        "//proto:remote_execution_go_proto",
        "//server/interfaces",
    ],
)

go_test(
    name = "bare_test",
    srcs = ["bare_test.go"],
    deps = [
        ":bare",
        "//enterprise/server/remote_execution/container",
        "//proto:remote_execution_go_proto",
        "//server/testutil/testfs",
        "@com_github_stretchr_testify//assert",
    ],
)
