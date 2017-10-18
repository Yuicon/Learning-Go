package main

import (
	"fmt"
)

func (s *IntSet) AddAll(vars ...int) {
	for _, x := range vars {
		s.Add(x)
	}
}

func main() {
	var x IntSet
	x.AddAll(1, 2, 3, 4)
	fmt.Println(x)
}