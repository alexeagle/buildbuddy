load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@bazel_tools//tools/build_defs/pkg:pkg.bzl", "pkg_tar")

genrule(
    name = "bazel-3.7_crossplatform",
    srcs = select({
        "@bazel_tools//src/conditions:darwin": ["@io_bazel_bazel-3.7-darwin-x86_64//file:downloaded"],
        "//conditions:default": ["@io_bazel_bazel-3.7-linux-x86_64//file:downloaded"],
    }),
    outs = ["bazel-3.7"],
    cmd_bash = "cp $(SRCS) $@",
    visibility = ["//visibility:public"],
)

genrule(
    name = "bazel-4.1_crossplatform",
    srcs = select({
        "@bazel_tools//src/conditions:darwin": ["@io_bazel_bazel-4.1-darwin-x86_64//file:downloaded"],
        "//conditions:default": ["@io_bazel_bazel-4.1-linux-x86_64//file:downloaded"],
    }),
    outs = ["bazel-4.1"],
    cmd_bash = "cp $(SRCS) $@",
    visibility = ["//visibility:public"],
)

genrule(
    name = "bazel-5.0.0rc3_crossplatform",
    srcs = select({
        "@bazel_tools//src/conditions:darwin": ["@io_bazel_bazel-5.0.0rc3-darwin-x86_64//file:downloaded"],
        "//conditions:default": ["@io_bazel_bazel-5.0.0rc3-linux-x86_64//file:downloaded"],
    }),
    outs = ["bazel-5.0.0rc3"],
    cmd_bash = "cp $(SRCS) $@",
    visibility = ["//visibility:public"],
)

pkg_tar(
    name = "bazel_binaries_tar",
    srcs = [
        ":bazel-3.7_crossplatform",
        ":bazel-4.1_crossplatform",
    ],
    package_dir = "/bazel",
    visibility = ["//visibility:public"],
)

go_library(
    name = "bazel",
    srcs = ["bazel.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/server/util/bazel",
    visibility = ["//visibility:public"],
    deps = [
        "//server/util/log",
        "//server/util/status",
    ],
)
