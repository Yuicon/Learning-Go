package main

import (
	"fmt"
)

func (s *IntSet) Elems() []int  {
	var tmp []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				tmp = append(tmp, 64*i+j)
			}
		}
	}
	return tmp
}

func main() {
	var x IntSet
	x.Add(16)
	x.Add(9)
	fmt.Println(x.Elems())
}