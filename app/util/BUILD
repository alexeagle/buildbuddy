load("//rules/typescript:index.bzl", "ts_jasmine_node_test", "ts_library")

package(default_visibility = ["//visibility:public"])

ts_library(
    name = "animated_value",
    srcs = ["animated_value.ts"],
    deps = [
        ":math",
        "//:node_modules/tslib",
    ],
)

ts_library(
    name = "animation_loop",
    srcs = ["animation_loop.ts"],
    deps = [
        ":time_delta",
        "//:node_modules/tslib",
    ],
)

ts_library(
    name = "async",
    srcs = ["async.ts"],
    deps = [
        "//:node_modules/tslib",
    ],
)

ts_jasmine_node_test(
    name = "async_test",
    srcs = ["async_test.ts"],
    deps = [
        ":async",
        "//:node_modules/tslib",
    ],
)

ts_library(
    name = "clipboard",
    srcs = ["clipboard.ts"],
)

ts_library(
    name = "dom",
    srcs = ["dom.ts"],
    deps = [
        "//:node_modules/tslib",
    ],
)

ts_library(
    name = "errors",
    srcs = ["errors.ts"],
)

ts_library(
    name = "git",
    srcs = ["git.ts"],
)

ts_jasmine_node_test(
    name = "git_test",
    srcs = ["git_test.ts"],
    deps = [
        ":git",
    ],
)

ts_library(
    name = "time_delta",
    srcs = ["time_delta.ts"],
    deps = [
        "//:node_modules/tslib",
    ],
)

ts_library(
    name = "math",
    srcs = ["math.ts"],
    deps = [
        "//:node_modules/tslib",
    ],
)

ts_library(
    name = "memo",
    srcs = ["memo.ts"],
    deps = [
        "//:node_modules/tslib",
    ],
)

ts_library(
    name = "proto",
    srcs = ["proto.ts"],
    deps = [
        "//proto:duration_ts_proto",
        "//proto:timestamp_ts_proto",
        "//:node_modules/@types/long",
        "//:node_modules/long",
    ],
)
