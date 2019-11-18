package stringsconcat

import "testing"

var input = []string {
	"ABDICATIVE", "ABDICATOR", "ABDICATORS" ,"ABDOMEN", "ABDOMENS" ,"ABDOMINA", "ABDOMINAL" ,"ABDOMINALLY" ,"ABDOMINALS" ,"ABDOMINOPLASTY" ,"ABDOMINOUS" ,"ABDUCE" ,"ABDUCED" ,"ABDUCENS" ,"ABDUCENT" ,"ABDUCENTES", "ABDUCES" ,"ABDUCING" ,"ABDUCT", "ABDUCTED", "ABDUCTEE", "ABDUCTEES" ,"ABDUCTING" ,"ABDUCTION" ,"ABDUCTIONS", "ABDUCTOR" ,"ABDUCTORES", "ABDUCTORS",
}

func BenchmarkConcatenation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatenation(input...)
	}
}

func BenchmarkBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		builder(input...)
	}
}

func BenchmarkBuilderWithPrealloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		builderWithPrealloc(input...)
	}
}
