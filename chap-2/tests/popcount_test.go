package popcount

import (
	"chap-2/popcount"
	// "fmt"
	"testing"
)

func BenchmarkPopCountSingleExpression(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(255)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountLoop(255)
	}
}
