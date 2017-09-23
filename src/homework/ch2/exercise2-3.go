package main

import (
	"fmt"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	fmt.Printf("%d\n", PopCount(999))
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var count byte
	for i := uint(0); i < 9; i++ {
		count += pc[byte(x>>(i*8))]
	}
	return int(count)
}
