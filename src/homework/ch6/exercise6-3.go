package main

import (
	"fmt"
)

func (s *IntSet) IntersectWith(t *IntSet) IntSet {
	var tmp IntSet
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 && t.Has(64*i+j) {
				tmp.Add(64*i+j)
			}
		}
	}
	return tmp
}

func (s *IntSet) DifferenceWith(t *IntSet) IntSet {
	var tmp IntSet
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 && !t.Has(64*i+j) {
				tmp.Add(64*i+j)
			}
		}
	}
	return tmp
}

func (s *IntSet) SymmetricDifference(t *IntSet) IntSet {
	var tmp IntSet
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 && !t.Has(64*i+j) {
				tmp.Add(64*i+j)
			}
		}
	}
	for i, word := range t.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 && !s.Has(64*i+j) {
				tmp.Add(64*i+j)
			}
		}
	}
	return tmp
}

func main() {
	var x, y IntSet
	x.Add(16)
	x.Add(9)

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"
	t := x.IntersectWith(&y)
	fmt.Println(t.String())
	d := x.DifferenceWith(&y)
	fmt.Println(d.String())
	s := x.SymmetricDifference(&y)
	fmt.Println(s.String())
}