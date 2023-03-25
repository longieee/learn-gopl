package popcount

import (
	"testing"
)

func BenchmarkPopCountSingleExpression(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(255)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(255)
	}
}
