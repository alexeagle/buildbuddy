load("//rules/typescript:index.bzl", "ts_library")

package(default_visibility = ["//visibility:public"])

exports_files(glob(["*.css"]))

ts_library(
    name = "digest",
    srcs = glob([
        "*.ts",
        "*.tsx",
    ]),
    deps = [
        "//app/format",
        "//:node_modules/@types/long",
        "//:node_modules/@types/react",
        "//:node_modules/long",
        "//:node_modules/react",
        "//:node_modules/tslib",
    ],
)
