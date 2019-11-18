#!/usr/bin/env bash

go test -benchmem -bench=. -p=1
#go test -benchmem -bench=BenchmarkConcatenation -p=1 -cpuprofile=concat.out
#go test -benchmem -bench=BenchmarkBuilder -p=1 -cpuprofile=builder.out
#go test -benchmem -bench=BenchmarkBuilderWithPrealloc -p=1 -cpuprofile=prealloc.out
rm *.test

#go tool pprof -cum -http=:8080 concat.out
#go tool pprof -cum -http=:8080 builder.out
#go tool pprof -cum -http=:8080 prealloc.out