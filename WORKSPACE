workspace(name = "protoconf")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

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

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "a82a352bffae6bee4e95f68a8d80a70e87f42c4741e6a448bec11998fcc82329",
    urls = [
        "https://github.com/bazelbuild/rules_go/releases/download/0.18.5/rules_go-0.18.5.tar.gz",
    ],
)

load(
    "@io_bazel_rules_go//go:deps.bzl",
    "go_register_toolchains",
    "go_rules_dependencies",
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "3c681998538231a2d24d0c07ed5a7658cb72bfb5fd4bf9911157c0e9ac6a2687",
    urls = [
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.17.0/bazel-gazelle-0.17.0.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

go_repository(
    name = "com_github_docker_libkv",
    commit = "458977154600b9f23984d9f4b82e79570b5ae12b",
    importpath = "github.com/docker/libkv",
    patches = ["//third_party:consul_fail_on_missing_key.patch"],
)

# Implicitly used by libkv
go_repository(
    name = "com_github_hashicorp_consul",
    importpath = "github.com/hashicorp/consul",
    tag = "v1.0.7",
)

# Implicitly used by libkv
go_repository(
    name = "com_github_samuel_go_zookeeper",
    commit = "c4fab1ac1bec58281ad0667dc3f0907a9476ac47",
    importpath = "github.com/samuel/go-zookeeper",
)

# Used by zookeeper
go_repository(
    name = "com_github_coreos_go_semver",
    importpath = "github.com/coreos/go-semver",
    tag = "v0.2.0",
)

go_repository(
    name = "com_github_jhump_protoreflect",
    commit = "6c4c7792338ef4769325550489b407691790ffa1",
    importpath = "github.com/jhump/protoreflect",
    patch_args = ["-p1"],
    patches = ["//third_party:protoreflect_proto_std_lib.patch"],
)

go_repository(
    name = "net_starlark_go",
    commit = "d6561f809f318cb4098a9e17073b3dfbf45d3289",
    importpath = "go.starlark.net",
)

# Used by starlark
go_repository(
    name = "com_github_chzyer_readline",
    commit = "2972be24d48e78746da79ba8e24e8b488c9880de",
    importpath = "github.com/chzyer/readline",
)

go_repository(
    name = "com_github_fsnotify_fsnotify",
    commit = "1d13583d846ea9d66dcabbfefbfb9d8e6fb05216",
    importpath = "github.com/fsnotify/fsnotify",
)

go_repository(
    name = "com_github_mitchellh_cli",
    importpath = "github.com/mitchellh/cli",
    tag = "v1.0.0",
)

# Used by cli
go_repository(
    name = "com_github_posener_complete",
    importpath = "github.com/posener/complete",
    tag = "v1.1.1",
)

# Used by cli
go_repository(
    name = "com_github_mattn_go_isatty",
    importpath = "github.com/mattn/go-isatty",
    tag = "v0.0.3",
)

# Used by cli
go_repository(
    name = "com_github_bgentry_speakeasy",
    importpath = "github.com/bgentry/speakeasy",
    tag = "v0.1.0",
)

# Used by cli
go_repository(
    name = "com_github_armon_go_radix",
    commit = "7fddfc383310abc091d79a27f116d30cf0424032",
    importpath = "github.com/armon/go-radix",
)

# Used by cli
go_repository(
    name = "com_github_fatih_color",
    importpath = "github.com/fatih/color",
    tag = "v1.7.0",
)

# Used by cli
go_repository(
    name = "com_github_hashicorp_go_multierror",
    importpath = "github.com/hashicorp/go-multierror",
    tag = "v1.0.0",
)

# Implicitly used by cli
go_repository(
    name = "com_github_hashicorp_errwrap",
    importpath = "github.com/hashicorp/errwrap",
    tag = "v1.0.0",
)

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

git_repository(
    name = "com_github_golang_protobuf",
    commit = "c823c79ea1570fb5ff454033735a8e68575d1d0f",  # v1.3.0, as of 2019-03-03
    patch_args = ["-p1"],
    patches = [
        "//third_party:com_github_golang_protobuf-gazelle.patch",
        "@io_bazel_rules_go//third_party:com_github_golang_protobuf-extras.patch",
        "//third_party:protobuf_fix_any_indentation.patch",
    ],
    remote = "https://github.com/golang/protobuf",
    shallow_since = "1549405252 -0800",
)
# gazelle args: -go_prefix github.com/golang/protobuf -proto disable_global

go_rules_dependencies()

go_register_toolchains(nogo = "@//:protoconf_nogo")

http_archive(
    name = "com_github_bazelbuild_buildtools",
    sha256 = "fabcd8a7f593f6dbe010fffb1d7e032438bd61342ccf0d4791e5211ea01e994a",
    strip_prefix = "buildtools-f720930ceb608b8c0d09528440ce1adeb01e61e0",
    urls = [
        # Master branch as of 2019-07-31
        "https://github.com/bazelbuild/buildtools/archive/f720930ceb608b8c0d09528440ce1adeb01e61e0.tar.gz",
    ],
)

load("@com_github_bazelbuild_buildtools//buildifier:deps.bzl", "buildifier_dependencies")

buildifier_dependencies()

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

load("//libprotoconf:deps.bzl", libprotoconf_deps = "deps")

libprotoconf_deps()

# No more `go_repository` behind this line, please move to appropriate deps.bzl file

go_repository(
    name = "com_github_mgutz_logxi",
    commit = "aebf8a7d67ab4625e0fd4a665766fef9a709161b",
    importpath = "github.com/mgutz/logxi",
)

go_repository(
    name = "com_github_hashicorp_go_uuid",
    commit = "4f571afc59f3043a65f8fe6bf46d887b10a01d43",
    importpath = "github.com/hashicorp/go-uuid",
)

go_repository(
    name = "com_github_mattn_go_colorable",
    commit = "388941e3ea99c461d2bc2747eaf27741b6dda096",
    importpath = "github.com/mattn/go-colorable",
)

go_repository(
    name = "com_github_mgutz_ansi",
    commit = "9520e82c474b0a04dd04f8a40959027271bab992",
    importpath = "github.com/mgutz/ansi",
)

go_repository(
    name = "org_golang_x_time",
    commit = "9d24e82272b4f38b78bc8cff74fa936d31ccd8ef",
    importpath = "golang.org/x/time",
)

go_repository(
    name = "com_github_hashicorp_go_rootcerts",
    commit = "df8e78a645e18d56ed7bb9ae10ffb8174ab892e2",
    importpath = "github.com/hashicorp/go-rootcerts",
)

go_repository(
    name = "in_gopkg_square_go_jose_v2",
    commit = "730df5f748271903322feb182be83b43ebbbe27d",
    importpath = "gopkg.in/square/go-jose.v2",
)

go_repository(
    name = "com_github_hashicorp_go_retryablehttp",
    commit = "6bb8533de58c768084f72589fcbb385719bc9dc5",
    importpath = "github.com/hashicorp/go-retryablehttp",
)

go_repository(
    name = "com_github_hashicorp_hcl",
    commit = "cf7d376da96d9cecec7c7483cec2735efe54a410",
    importpath = "github.com/hashicorp/hcl",
)

go_repository(
    name = "com_github_hashicorp_go_cleanhttp",
    commit = "d3fcbee8e1810ecee4bdbf415f42f84cfd0e3361",
    importpath = "github.com/hashicorp/go-cleanhttp",
)

go_repository(
    name = "com_github_mitchellh_mapstructure",
    commit = "3536a929edddb9a5b34bd6861dc4a9647cb459fe",
    importpath = "github.com/mitchellh/mapstructure",
)

go_repository(
    name = "com_github_hashicorp_go_hclog",
    commit = "f1d61ad5398ffe4f2eb61eacb088340d44e99672",
    importpath = "github.com/hashicorp/go-hclog",
)

go_repository(
    name = "com_github_elazarl_go_bindata_assetfs",
    commit = "38087fe4dafb822e541b3f7955075cc1c30bd294",
    importpath = "github.com/elazarl/go-bindata-assetfs",
)

go_repository(
    name = "com_github_jefferai_jsonx",
    commit = "a29fe22bc88c13040654d8381139d2433a5b76b6",
    importpath = "github.com/jefferai/jsonx",
)

go_repository(
    name = "com_github_mitchellh_go_testing_interface",
    commit = "6d0b8010fcc857872e42fc6c931227569016843c",
    importpath = "github.com/mitchellh/go-testing-interface",
)

go_repository(
    name = "com_github_mitchellh_copystructure",
    commit = "9a1b6f44e8da0e0e374624fb0a825a231b00c537",
    importpath = "github.com/mitchellh/copystructure",
)

go_repository(
    name = "com_github_nytimes_gziphandler",
    commit = "dd0439581c7657cb652dfe5c71d7d48baf39541d",
    importpath = "github.com/nytimes/gziphandler",
)

go_repository(
    name = "com_github_hashicorp_go_sockaddr",
    commit = "c7188e74f6acae5a989bdc959aa779f8b9f42faf",
    importpath = "github.com/hashicorp/go-sockaddr",
)

go_repository(
    name = "com_github_mitchellh_reflectwalk",
    commit = "3e2c75dfad4fbf904b58782a80fd595c760ad185",
    importpath = "github.com/mitchellh/reflectwalk",
)

go_repository(
    name = "com_github_keybase_go_crypto",
    commit = "b785b22cc75714432d7ebf0b554dfafaff608671",
    importpath = "github.com/keybase/go-crypto",
)

go_repository(
    name = "org_golang_x_crypto",
    commit = "4def268fd1a49955bfb3dda92fe3db4f924f2285",
    importpath = "golang.org/x/crypto",
)

go_repository(
    name = "com_github_patrickmn_go_cache",
    commit = "5633e0862627c011927fa39556acae8b1f1df58a",
    importpath = "github.com/patrickmn/go-cache",
)

go_repository(
    name = "com_github_armon_go_metrics",
    commit = "ec5e00d3c878b2a97bbe0884ef45ffd1b4f669f5",
    importpath = "github.com/armon/go-metrics",
)

go_repository(
    name = "com_github_hashicorp_golang_lru",
    commit = "7f827b33c0f158ec5dfbba01bb0b14a4541fd81d",
    importpath = "github.com/hashicorp/golang-lru",
)

go_repository(
    name = "com_github_oklog_run",
    commit = "6934b124db28979da51d3470dadfa34d73d72652",
    importpath = "github.com/oklog/run",
)

go_repository(
    name = "com_github_hashicorp_go_memdb",
    commit = "5500ca0de0dab231b02aedabac095d43a59f31d2",
    importpath = "github.com/hashicorp/go-memdb",
)

go_repository(
    name = "com_github_ryanuber_go_glob",
    commit = "51a8f68e6c24dc43f1e371749c89a267de4ebc53",
    importpath = "github.com/ryanuber/go-glob",
)

go_repository(
    name = "com_github_hashicorp_go_immutable_radix",
    commit = "0146a9aba1948ded4ed290cfd3fded2c15313f63",
    importpath = "github.com/hashicorp/go-immutable-radix",
)

go_repository(
    name = "com_github_pierrec_lz4",
    commit = "a54ef8c8617bd3d46bed8dd6c0d5547f06093d30",
    importpath = "github.com/pierrec/lz4",
)

go_repository(
    name = "com_github_golang_snappy",
    commit = "2a8bb927dd31d8daada140a5d09578521ce5c36a",
    importpath = "github.com/golang/snappy",
)

go_repository(
    name = "com_github_jeffail_gabs",
    commit = "829ed3fc9fc24d2f0bbf7af14b40c8a399196bd1",
    importpath = "github.com/jeffail/gabs/v2",
)

go_repository(
    name = "com_github_hashicorp_go_plugin",
    commit = "9e3e1c37db188a1acb66561ee0ed4bf4d5e77554",
    importpath = "github.com/hashicorp/go-plugin",
)

go_repository(
    name = "com_github_hashicorp_go_version",
    commit = "192140e6f3e645d971b134d4e35b5191adb9dfd3",
    importpath = "github.com/hashicorp/go-version",
)
