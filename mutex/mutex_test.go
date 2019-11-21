package mutex

import (
	"math/rand"
	"testing"
)

func BenchmarkSingletonUse(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rand.Int()
		}
	})
}


func BenchmarkConcurrentUse(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		myInstance := rand.New(rand.NewSource(42))
		for pb.Next() {
			myInstance.Int()
		}
	})
}


