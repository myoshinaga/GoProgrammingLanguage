package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCountTable(x uint64) int {
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
	count := 0
	for i := 0; i < 8; i++ {
		count += int(pc[byte(x>>(uint(i)*8))])
	}
	return count
}

func PopCountLowestBit(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		count += int((x >> uint(i)) & uint64(1))
	}
	return count
}

func PopCountLowestBitClear(x uint64) int {
	count := 0
	for ; x != 0; count++ {
		x &= x - 1
	}
	return count
}
