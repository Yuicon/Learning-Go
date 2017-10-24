package main

import (
	"fmt"
	"bufio"
	"strconv"
)

type Counter struct {
	word, line int
	p []byte
}

func (c *Counter) Write(p []byte) (int, error) {
	c.word = split(p, bufio.ScanWords)
	c.line = split(p, bufio.ScanLines)
	return len(p), nil
}

func (c Counter) String() string {
	return "word:" + strconv.Itoa(c.word) + ",line:" + strconv.Itoa(c.line)
}

func split(p []byte, s bufio.SplitFunc) int {
	if len(p) == 0 {
		return 0
	}
	result := 1
	advance, _, _ := s(p, false)
	end := advance
	for advance > 0 {
		if end >= len(p) {
			break
		}
		advance, _, _ = s(p[end:], false)
		end += advance
		result++
	}
	return result
}

func main() {
	var c Counter
	p := []byte("Build encapsulated components that manage their own state,\n then compose them to make complex UIs.")
	fmt.Println(split(p, bufio.ScanWords))
	fmt.Println(split(p, bufio.ScanLines))
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")
}
