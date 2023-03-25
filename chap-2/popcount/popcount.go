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

func PopCountLoop(x uint64) int {
	var ret int = 0
	for i := 0; i < 8; i++ {
		ret += int(pc[byte(x>>(i*8))])
	}
	return ret
}
