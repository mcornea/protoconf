load("@bazel_gazelle//:deps.bzl", "go_repository")

def deps():
    go_repository(
        name = "com_github_hashicorp_vault",
        tag = "v1.2.2",
        importpath = "github.com/hashicorp/vault",
        # build_extra_args = ["-exclude=vendor"],
    )

    go_repository(
        name = "com_github_jcmturner_vaultmock",
        commit = "348b994639e19b2178c7d6c1c1da1a95f08adee0",
        importpath = "github.com/jcmturner/vaultmock",
        build_extra_args = ["-exclude=vendor"],
    )
