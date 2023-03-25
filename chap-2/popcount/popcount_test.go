package popcount

import (
	"testing"
)

func BenchmarkPopCountSingleExpression(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0xFFFFFFFFFFFFFFFF)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(0xFFFFFFFFFFFFFFFF)
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(0xFFFFFFFFFFFFFFFF)
	}
}

func TestPopCountLoop(t *testing.T) {
	tcs := []struct {
		number  uint64
		expects int
	}{
		{0x1234567890ABCDEF, 32},
		{0xFFFFFFFFFFFFFFFF, 64},
		{0x0000000000000002, 1},
		{0x0000000000000000, 0},
		{0x1000000000000000, 1},
	}

	for _, tc := range tcs {
		ret := PopCountLoop(tc.number)
		if ret != tc.expects {
			t.Errorf("PopCountLoop Failed. Number: %X, expect counts: %d, get: %d", tc.number, tc.expects, ret)
		}
	}
}

func TestPopCountShift(t *testing.T) {
	tcs := []struct {
		number  uint64
		expects int
	}{
		{0x1234567890ABCDEF, 32},
		{0xFFFFFFFFFFFFFFFF, 64},
		{0x0000000000000002, 1},
		{0x0000000000000000, 0},
		{0x1000000000000000, 1},
	}

	for _, tc := range tcs {
		ret := PopCountShift(tc.number)
		if ret != tc.expects {
			t.Errorf("PopCountShift Failed. Number: %X, expect counts: %d, get: %d", tc.number, tc.expects, ret)
		}
	}
}
