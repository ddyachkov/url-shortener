package app

import (
	"math/rand"
	"testing"
)

func Benchmark_makeURI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		makeURI(rand.Intn(1000))
	}
}
