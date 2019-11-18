package regex

import "testing"

var input = []string {
	"ABDICATIVE", "ABDICATOR", "ABDICATORS" ,"ABDOMEN", "ABDOMENS" ,"ABDOMINA", "ABDOMINAL" ,"ABDOMINALLY" ,"ABDOMINALS" ,"ABDOMINOPLASTY" ,"ABDOMINOUS" ,"ABDUCE" ,"ABDUCED" ,"ABDUCENS" ,"ABDUCENT" ,"ABDUCENTES", "ABDUCES" ,"ABDUCING" ,"ABDUCT", "ABDUCTED", "ABDUCTEE", "ABDUCTEES" ,"ABDUCTING" ,"ABDUCTION" ,"ABDUCTIONS", "ABDUCTOR" ,"ABDUCTORES", "ABDUCTORS",
}

func BenchmarkNormal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := range input {
			normal(input[j])
		}
	}
}


func BenchmarkCached(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := range input {
			cached(input[j])
		}
	}
}
