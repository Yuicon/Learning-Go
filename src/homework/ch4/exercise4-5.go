package main

import (
	"fmt"
)

func main() {
	a := []string{"m", "m", "a", "a", "a", "x", "x"}
	a = noRepetition(a)
	fmt.Println(a)
}

func noRepetition(s []string) []string {
	var a []string
	var tmp string
	for _, t := range s {
		if tmp != t {
			a = append(a, t)
		}
		tmp = t
	}
	return a
}
