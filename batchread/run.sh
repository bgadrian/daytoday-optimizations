#!/usr/bin/env bash

go test -benchmem -bench=BenchmarkGetRandValue -cpuprofile=simple.out
go test -benchmem -bench=BenchmarkSimpleList_GetRandValue -cpuprofile=batch.out
rm *.test

go tool pprof -cum -http=:8080 simple.out
go tool pprof -cum -http=:8080 batch.out