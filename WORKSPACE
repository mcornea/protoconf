workspace(name = "protoconf")

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "com_github_golang_protobuf",
    urls = [
        "https://github.com/golang/protobuf/archive/v1.3.1.tar.gz",
    ],
    strip_prefix = "protobuf-1.3.1",
    patch_args = ["-p1"],
    patches = [
        "@//third_party:com_github_golang_protobuf-gazelle.patch",
        "@io_bazel_rules_go//third_party:com_github_golang_protobuf-extras.patch",
        "@//third_party:protobuf_fix_any_indentation.patch",
    ],
    sha256 = "3f3a6123054a9847093c119895f1660612f301fe95358f3a6a1a33fd0933e6cf",
)
# gazelle args: -go_prefix github.com/golang/protobuf -proto disable_global

http_archive(
    name = "io_bazel_rules_go",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/rules_go/releases/download/v0.20.3/rules_go-v0.20.3.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.20.3/rules_go-v0.20.3.tar.gz",
    ],
    sha256 = "e88471aea3a3a4f19ec1310a55ba94772d087e9ce46e41ae38ecebe17935de7b",
)

http_archive(
    name = "bazel_gazelle",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/bazel-gazelle/releases/download/v0.19.1/bazel-gazelle-v0.19.1.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.19.1/bazel-gazelle-v0.19.1.tar.gz",
    ],
    sha256 = "86c6d481b3f7aedc1d60c1c211c6f76da282ae197c3b3160f54bd3a8f847896f",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(nogo = "@//:protoconf_nogo")

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()

http_archive(
    name = "bazel_skylib",
    sha256 = "1dde365491125a3db70731e25658dfdd3bc5dbdfd11b840b3e987ecf043c7ca0",
    url = "https://github.com/bazelbuild/bazel-skylib/releases/download/0.9.0/bazel_skylib-0.9.0.tar.gz",
)

load("@bazel_skylib//:workspace.bzl", "bazel_skylib_workspace")

bazel_skylib_workspace()

http_archive(
    name = "rules_proto",
    sha256 = "602e7161d9195e50246177e7c55b2f39950a9cf7366f74ed5f22fd45750cd208",
    strip_prefix = "rules_proto-97d8af4dc474595af3900dd85cb3a29ad28cc313",
    urls = [
        # Master branch as of 2019-08-01
        "https://mirror.bazel.build/github.com/bazelbuild/rules_proto/archive/97d8af4dc474595af3900dd85cb3a29ad28cc313.tar.gz",
        "https://github.com/bazelbuild/rules_proto/archive/97d8af4dc474595af3900dd85cb3a29ad28cc313.tar.gz",
    ],
)

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")

rules_proto_dependencies()

rules_proto_toolchains()

# Docker
http_archive(
    name = "containerregistry",
    patch_args = ["-p1"],
    patches = ["//third_party:containerregistry.patch"],
    sha256 = "a8cdf2452323e0fefa4edb01c08b2ec438c9fa3192bc9f408b89287598c12abc",
    strip_prefix = "containerregistry-0.0.36",
    urls = ["https://github.com/google/containerregistry/archive/v0.0.36.tar.gz"],
)

http_archive(
    name = "io_bazel_rules_docker",
    patch_args = ["-p1"],
    patches = ["//third_party:rules_docker_container_push.patch"],
    sha256 = "87fc6a2b128147a0a3039a2fd0b53cc1f2ed5adb8716f50756544a572999ae9a",
    strip_prefix = "rules_docker-0.8.1",
    urls = ["https://github.com/bazelbuild/rules_docker/archive/v0.8.1.tar.gz"],
)

load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)

container_repositories()

load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()

load("@com_github_bazelbuild_buildtools//buildifier:deps.bzl", "buildifier_dependencies")

buildifier_dependencies()

load("@//:deps.bzl", "deps")

deps()

load("@//pc4tf:deps.bzl", pc4tf_deps="deps")

pc4tf_deps()
