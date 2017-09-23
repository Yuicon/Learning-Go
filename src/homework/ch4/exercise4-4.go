package main

import (
	"fmt"
)

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6}
	a = rotate(a, 3)
	fmt.Println(a)
}

// reverse reverses a slice of ints in place.
func rotate(s []int, n int) []int {
	tmp := s[n:]
	for i := 0; i < n; i++ {
		tmp = append(tmp, s[i])
	}
	return tmp
}
