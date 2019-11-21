#!/usr/bin/env bash

#go test -benchmem -bench=.
go test -benchmem -bench=BenchmarkSingletonUse -cpuprofile=normal.out
go test -benchmem -bench=BenchmarkConcurrentUse -cpuprofile=concurrent.out
rm *.test

go tool pprof -cum -http=:8080 normal.out
#go tool pprof -cum -http=:8080 concurrent.out