# kopsexample

```
$ go run pkg/main.go
cluster {{ } {      0 0001-01-01 00:00:00 +0000 UTC <nil> <nil> map[] map[] [] []  []} { []   <nil>   []     []  <nil>     <nil> []     [] [] <nil> <nil> [] <nil> <nil> <nil> <nil> [] [] <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> map[] [] <nil> <nil> <nil> false <nil> <nil> [] <nil>}}. version 1.11.0
```

```
$ bazel run //pkg:pkg
ERROR: /private/var/tmp/_bazel_odashevskii/94eefb94fd8aa518c28b0c8c0ea6d712/external/io_k8s_kops/upup/models/BUILD.bazel:14:8: no such package '@io_k8s_kops//vendor/github.com/go-bindata/go-bindata/go-bindata': BUILD file not found in directory 'vendor/github.com/go-bindata/go-bindata/go-bindata' of external repository @io_k8s_kops. Add a BUILD file to a directory to mark it as a package. and referenced by '@io_k8s_kops//upup/models:bindata'
ERROR: Analysis of target '//pkg:pkg' failed; build aborted: Analysis failed
INFO: Elapsed time: 3.058s
INFO: 0 processes.
FAILED: Build did NOT complete successfully (4 packages loaded, 166 targets configured)
FAILED: Build did NOT complete successfully (4 packages loaded, 166 targets configured)
```

The culprit is https://github.com/kubernetes/kops/blob/master/upup/models/BUILD.bazel:

```bazel
load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "bindata.go",
        "vfs.go",
    ],
    importpath = "k8s.io/kops/upup/models",
    visibility = ["//visibility:public"],
    deps = ["//util/pkg/vfs:go_default_library"],
)

genrule(
    name = "bindata",
    srcs = glob(
        [
            "cloudup/**",
            "nodeup/**",
        ],
    ),
    outs = ["bindata.go"],
    cmd = """
$(location //vendor/github.com/go-bindata/go-bindata/v3/go-bindata:go-bindata) \
  -o "$(OUTS)" -pkg models \
  -nometadata \
  -nocompress \
  -prefix $$(pwd) \
  -prefix upup/models $(SRCS)
""",
    tools = [
        "//vendor/github.com/go-bindata/go-bindata/v3/go-bindata",
    ],
)
```
