# go_bench_tests

## run
Command:
```
go test --bench=. -test.benchmem -cpuprofile profile_cpu.out -memprofile profile_mem.out .
```

Result:
```
goos: darwin
goarch: amd64
pkg: laullon.com/bench_tests
BenchmarkQueueFix-8     	 5000000	       312 ns/op	      32 B/op	       1 allocs/op
BenchmarkQueueList-8    	 5000000	       380 ns/op	      80 B/op	       2 allocs/op
BenchmarkQueueSlice-8   	 5000000	       345 ns/op	      48 B/op	       2 allocs/op
PASS
ok  	laullon.com/bench_tests	6.425s
```

## Generate pdf's

```
go tool pprof -pdf profile_cpu.out > profile_cpu.pdf
```

[profile_cpu.pdf](https://github.com/laullon/go_bench_tests/blob/master/profile_cpu.pdf)

```
go tool pprof -pdf profile_mem.out > profile_mem.pdf
```

[profile_mem.pdf](https://github.com/laullon/go_bench_tests/blob/master/profile_mem.pdf)
