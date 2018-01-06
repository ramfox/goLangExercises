// Go Programming Language Book
//
// example 4.1
//
// sha265Dif takes the sha256 hash of two inputs and counts
// the number of bits that are different

package sha256Diff

import ()

// pc[i] is the population count of i.
var pc [256]byte

type hash [32]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x hash) int {
	var popCount int
	for i := 0; i < 32; i++ {
		popCount += int(pc[x[i]])
	}
	return popCount
}

// Diff returns the number of bits different in two hashes ([32]byte)
func Diff(a, b hash) int {
	var x hash
	for i := 0; i < 32; i++ {
		x[i] = a[i] ^ b[i]
	}
	return PopCount(x)
}
