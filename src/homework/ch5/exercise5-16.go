package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%s\n", join(" | ", "321", "456", "789"))
}

func join(sep string, vals... string) string {
	switch len(vals) {
	case 0:
		return ""
	case 1:
		return vals[0]
	case 2:
		// Special case for common small values.
		// Remove if golang.org/issue/6714 is fixed
		return vals[0] + sep + vals[1]
	case 3:
		// Special case for common small values.
		// Remove if golang.org/issue/6714 is fixed
		return vals[0] + sep + vals[1] + sep + vals[2]
	}
	n := len(sep) * (len(vals) - 1)
	for i := 0; i < len(vals); i++ {
		n += len(vals[i])
	}

	b := make([]byte, n)
	bp := copy(b, vals[0])
	for _, s := range vals[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}
	return string(b)
}
