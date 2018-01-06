// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)
//!+
package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCountExpression returns the population count (number of set bits) of x.
func PopCountExpression(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// Ex 02.03
// PopCountLoop returns the population count (number of set bits) of x.
func PopCountLoop(x uint64) int {
	var popCount byte
	for i := uint(0); i < 64; i += 8 {
		popCount += pc[byte(x>>i)]
	}
	return int(popCount)
}

// Ex 02.04
// PopCountLoop returns the population count (number of set bits) of x.
func PopCountShift(x uint64) int {
	var popCount uint64
	for x != 0 {
		popCount += x & 1
		x = x >> 1
	}
	return int(popCount)
}

// Ex 02.05
func PopCountClear(x uint64) int {
	var popCount int
	for x != 0 {
		x = x & (x - 1)
		popCount++
	}
	return popCount
}

//!-
