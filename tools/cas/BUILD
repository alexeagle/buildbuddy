load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "cas_lib",
    srcs = ["cas.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/tools/cas",
    visibility = ["//visibility:private"],
    deps = [
        "//proto:remote_execution_go_proto",
        "//server/remote_cache/cachetools",
        "//server/remote_cache/digest",
        "//server/util/grpc_client",
        "//server/util/log",
        "//server/util/status",
        "@go_googleapis//google/bytestream:bytestream_go_proto",
        "@org_golang_google_grpc//metadata",
        "@org_golang_google_protobuf//encoding/prototext",
        "@org_golang_google_protobuf//proto",
    ],
)

go_binary(
    name = "cas",
    embed = [":cas_lib"],
    visibility = ["//visibility:public"],
)
