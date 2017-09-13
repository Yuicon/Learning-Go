package main

import (
	"fmt"
)

func main() {
	a := []string{"m", "a", "a", "x"}
	a = noRepetition(a)
	fmt.Println(a)
}

func noRepetition(s []string) []string {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			for j := i + 1; j < len(s)-1; j++ {
				s[j] = s[j+1]
				s = s[:len(s)-1]
			}
		}
	}
	return s
}
