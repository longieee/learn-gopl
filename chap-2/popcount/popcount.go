package popcount

// pc[i] is the population count of i.
// Poppulation count: the numberofset bits, that is, bits whose value is1, in a uint64 value
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// Ex 2.3
// Rewrite PopCount to use a loop instead of a single expression.
func PopCountLoop(x uint64) int {
	var ret int = 0
	for i := 0; i < 8; i++ {
		ret += int(pc[byte(x>>(i*8))])
	}
	return ret
}

// Ex 2.4
// Write a version of PopCount that counts bits by shifting its argument through 64
// bits position, testing the rightmost bit at each time.

func PopCountShift(x uint64) int {
	n := 0
	for i := uint64(0); i < 64; i++ {
		if x&1 != 0 {
			n++
		}
		x = x >> 1
	}
	return n
}
