Go Plugin Benchmarks
==

This is a benchmark script that compares the performance of the following methods as Go's plugin mechanism.

- Shared libraries (ex: standard library)
- External APIs (ex: hashicorp/go-plugin, knqyf263/go-plugin)
- External commands (ex: kubectl)

Usage
--

```
# make
go build -buildmode=plugin -o plugin.so ./so/main.go
go build -o plugin.cli cli/main.go
go test -bench . -benchmem
goos: linux
goarch: arm64
pkg: github.com/linyows/benchmark-plugins
BenchmarkSO-4              21206             54535 ns/op            5059 B/op         62 allocs/op
BenchmarkAPI-4              4578            294199 ns/op           23206 B/op        215 allocs/op
BenchmarkCLI-4               746           1603926 ns/op            8337 B/op         92 allocs/op
PASS
ok      github.com/linyows/benchmark-plugins    4.478s
```

Author
--

[linyows](https://github.com/linyows)
