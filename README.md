# kopsexample

```
$ bazel run //pkg:pkg                                                                                        6s 86ms
ERROR: /private/var/tmp/_bazel_odashevskii/94eefb94fd8aa518c28b0c8c0ea6d712/external/io_k8s_kops/pkg/apis/kops/BUILD.bazel:3:11: no such package '@io_k8s_kops//vendor/github.com/blang/semver': BUILD file not found in directory 'vendor/github.com/blang/semver' of external repository @io_k8s_kops. Add a BUILD file to a directory to mark it as a package. and referenced by '@io_k8s_kops//pkg/apis/kops:go_default_library'
ERROR: Analysis of target '//pkg:pkg' failed; build aborted: Analysis failed
INFO: Elapsed time: 0.102s
INFO: 0 processes.
FAILED: Build did NOT complete successfully (0 packages loaded, 0 targets configured)
FAILED: Build did NOT complete successfully (0 packages loaded, 0 targets configured)
```

