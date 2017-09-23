package main

import (
	"fmt"
	"bytes"
	"strings"
)

func main() {
	fmt.Printf("%s\n", comma2("122132134.213131"))
}

// comma inserts commas in a non-negative decimal integer string.
func comma2(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	var decimals string
	breakpoint := strings.Index(s, ".")
	if (breakpoint != -1) {
		decimals = s[breakpoint:]
		s = s[:breakpoint]
		n = len(s)
	}
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
	if (breakpoint != -1) {
		buf.WriteString(decimals)
	}
	return buf.String()
}
