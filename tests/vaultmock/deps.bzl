load("@bazel_gazelle//:deps.bzl", "go_repository")

def deps():
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
        importpath = "github.com/NYTimes/gziphandler",
        sum = "h1:ZUDjpQae29j0ryrS0u/B8HZfJBtBQHjqw2rQ2cqUQ3I=",
        version = "v1.1.1",
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
        name = "com_github_hashicorp_go_plugin",
        commit = "9e3e1c37db188a1acb66561ee0ed4bf4d5e77554",
        importpath = "github.com/hashicorp/go-plugin",
    )

    go_repository(
        name = "com_github_hashicorp_go_version",
        commit = "192140e6f3e645d971b134d4e35b5191adb9dfd3",
        importpath = "github.com/hashicorp/go-version",
    )

    go_repository(
        name = "com_github_hashicorp_yamux",
        commit = "2f1d1f20f75d5404f53b9edf6b53ed5505508675",
        importpath = "github.com/hashicorp/yamux",
    )

    go_repository(
        name = "com_github_prometheus_client_golang",
        commit = "35ef65db672a76effef5f0808decd0484a636f3f",
        importpath = "github.com/prometheus/client_golang",
    )

    go_repository(
        name = "com_github_hashicorp_raft_snapshot",
        commit = "a92b38874ad4515f642bb7c889274b1d42ae14b5",
        importpath = "github.com/hashicorp/raft-snapshot",
    )

    go_repository(
        name = "com_github_lib_pq",
        commit = "78223426e7c66d631117c0a9da1b7f3fde4d23a5",
        importpath = "github.com/lib/pq",
    )

    go_repository(
        name = "com_github_prometheus_common",
        commit = "637d7c34db122e2d1a25d061423098663758d2d3",
        importpath = "github.com/prometheus/common",
    )

    go_repository(
        name = "com_github_go_sql_driver_mysql",
        commit = "877a9775f06853f611fb2d4e817d92479242d1cd",
        importpath = "github.com/go-sql-driver/mysql",
    )

    go_repository(
        name = "io_etcd_go_bbolt",
        commit = "35b666109344c1547e0135d0d17486746cc073f3",
        importpath = "go.etcd.io/bbolt",
    )

    go_repository(
        name = "com_github_hashicorp_go_raftchunking",
        commit = "09594bcb5a8c0dafcf099cc621fc40ce54adce57",
        importpath = "github.com/hashicorp/go-raftchunking",
    )

    go_repository(
        name = "com_github_hashicorp_raft",
        commit = "db5ceea63eeda4e4022acfc1a333cc835ab957ac",
        importpath = "github.com/hashicorp/raft",
    )

    go_repository(
        name = "com_github_jeffail_gabs_v2",
        # commit = "829ed3fc9fc24d2f0bbf7af14b40c8a399196bd1",
        importpath = "github.com/Jeffail/gabs/v2",
        version = "v2.1.0",
        sum = "h1:6dV9GGOjoQgzWTQEltZPXlJdFloxvIq7DwqgxMCbq30=",
    )

    go_repository(
        name = "org_golang_x_net",
        importpath = "golang.org/x/net",
        sum = "h1:fHDIZ2oxGnUZRN6WgWFCbYBjH9uqVPRCUVUDhs0wnbA=",
        version = "v0.0.0-20190813141303-74dc4d7220e7",
    )

    go_repository(
        name = "org_golang_google_grpc",
        importpath = "google.golang.org/grpc",
        sum = "h1:AzbTB6ux+okLTzP8Ru1Xs41C303zdcfEht7MQnYJt5A=",
        version = "v1.23.0",
    )

    go_repository(
        name = "org_golang_x_text",
        importpath = "golang.org/x/text",
        sum = "h1:tW2bmiBqwgJj/UpqtC8EpXEZVYOwU0yG4iWbprSVAcs=",
        version = "v0.3.2",
    )

    go_repository(
        name = "com_github_hashicorp_go_msgpack",
        importpath = "github.com/hashicorp/go-msgpack",
        sum = "h1:i9R9JSrqIz0QVLz3sz+i3YJdT7TTSLcfLLzJi9aZTuI=",
        version = "v0.5.5",
    )

    go_repository(
        name = "com_github_matttproud_golang_protobuf_extensions",
        importpath = "github.com/matttproud/golang_protobuf_extensions",
        sum = "h1:4hp9jkHxhMHkqkrB3Ix0jegS5sx/RkqARlsWZ6pIwiU=",
        version = "v1.0.1",
    )

    go_repository(
        name = "com_github_prometheus_client_model",
        importpath = "github.com/prometheus/client_model",
        sum = "h1:gQz4mCbXsO+nc9n1hCxHcGA3Zx3Eo+UHZoInFGUIXNM=",
        version = "v0.0.0-20190812154241-14fe0d1b01d4",
    )

    go_repository(
        name = "com_github_prometheus_procfs",
        importpath = "github.com/prometheus/procfs",
        sum = "h1:CTwfnzjQ+8dS6MhHHu4YswVAD99sL2wjPqP+VkURmKE=",
        version = "v0.0.3",
    )

    go_repository(
        name = "com_github_beorn7_perks",
        importpath = "github.com/beorn7/perks",
        sum = "h1:VlbKKnNfV8bJzeqoa4cOKqO6bYr3WgKZxO8Z16+hsOM=",
        version = "v1.0.1",
    )
