package main

import (
	"fmt"
	"bytes"
)

func main() {
	fmt.Printf("%s\n", comma("1234"))
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	times := n - n / 3 * 3
	if times != 0 {
		buf.WriteString(s[:times])
		buf.WriteString(",")
	}
	for i := 0;i < n / 3; i++ {
		buf.WriteString(s[times + i * 3:times + i * 3 + 3])
		if times + i * 3 + 3 != n {
			buf.WriteString(",")
		}
	}
	return buf.String()
}
