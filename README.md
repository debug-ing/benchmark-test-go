# Benchmark Testing in go

A simple code and ci for benchmark test in golang

## Run
without benchmem
```bash
go test -bench=.
```

with benchmem
```bash
go test -bench=. -benchmem
```