workspace(name = "org_ckuenzl_go_api")

load("//:bazel/deps.bzl", "go_ckuenzl_dependencies")

go_ckuenzl_dependencies()

##############################################################################
# Protobuf
##############################################################################
load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")

rules_proto_dependencies()

rules_proto_toolchains()

##############################################################################
# gRPC
##############################################################################
load("@com_github_grpc_grpc//bazel:grpc_deps.bzl", "grpc_deps")

grpc_deps()

##############################################################################
# Golang
##############################################################################
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(
    go_version = "1.13.7",
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")


load("//:bazel/repositories.bzl", "go_ckuenzl_repositories")
# gazelle:repo bazel_gazelle
# gazelle:repository_macro bazel/repositories.bzl%go_ckuenzl_repositories
go_ckuenzl_repositories()