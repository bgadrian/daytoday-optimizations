package batchread

import (
	"math/rand"
	"testing"
)

const batchSize = 20

func BenchmarkGetRandValue(b *testing.B) {
	random := rand.New(rand.NewSource(233))

	for n := 0; n < b.N; n++ {
		for i := 0; i < batchSize; i++ {
			GetRandValue(random, "hacker", "phrase")
		}
	}
}

func BenchmarkSimpleList_GetRandValue(b *testing.B) {
	random := rand.New(rand.NewSource(233))
	shortcut, _ := NewDataListCache(random, "hacker", "phrase")

	for n := 0; n < b.N; n++ {
		for i := 0; i < batchSize; i++ {
			shortcut.GetRandValue()
		}
	}
}
