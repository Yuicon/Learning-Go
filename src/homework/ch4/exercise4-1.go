package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

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

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	var count int
	for i := range c1 {
		//fmt.Printf("%08b\n", c1[i])
		//fmt.Printf("%08b\n", c2[i])
		//fmt.Printf("%08b\n", c1[i]^c2[i])
		count += PopCount(uint64(c1[i] ^ c2[i]))
		//fmt.Printf("%d\n \n", count)
	}
	fmt.Printf("%d\n", count)
}
