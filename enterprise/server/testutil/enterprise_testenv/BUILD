load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "enterprise_testenv",
    testonly = 1,
    srcs = ["enterprise_testenv.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/testutil/enterprise_testenv",
    visibility = ["//visibility:public"],
    deps = [
        "//enterprise/server/backends/authdb",
        "//enterprise/server/backends/redis_cache",
        "//enterprise/server/backends/userdb",
        "//enterprise/server/util/redisutil",
        "//server/testutil/healthcheck",
        "//server/testutil/testenv",
        "//server/util/log",
        "//server/util/testing/flags",
        "@com_github_stretchr_testify//assert",
    ],
)
