#!/usr/bin/env bash

go test -benchmem -bench=.
go test -benchmem -bench=BenchmarkNormal -cpuprofile=normal.out
go test -benchmem -bench=BenchmarkCached -cpuprofile=cached.out
rm *.test

go tool pprof -cum -http=:8080 normal.out
go tool pprof -cum -http=:8080 cached.out