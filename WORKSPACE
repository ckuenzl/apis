workspace(name = "org_ckuenzl_go_api")

load("//:BAZEL/deps.bzl", "go_ckuenzl_dependencies")

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
# BAZEL
##############################################################################
load("@upb//bazel:repository_defs.bzl", "bazel_version_repository")

bazel_version_repository(
    name = "bazel_version",
)

##############################################################################
# Golang
##############################################################################
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(
    go_version = "1.15.3",
)

##############################################################################
# Gazelle
##############################################################################
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()


load("//:BAZEL/repositories.bzl", "go_ckuenzl_repositories", "go_ckuenzl_grpc_dependencies")

go_ckuenzl_grpc_dependencies()

# gazelle:repo bazel_gazelle
# gazelle:repository_macro BAZEL/repositories.bzl%go_ckuenzl_repositories
go_ckuenzl_repositories()