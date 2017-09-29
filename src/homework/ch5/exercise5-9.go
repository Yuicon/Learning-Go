package main

import (
	"strings"
	"fmt"
)

func main() {
	fmt.Printf("%s", expand("asfooodfooafoo", replace))
}

func expand(s string, f func(string) string) string {
	return strings.Replace(s, "foo", f("foo"), -1)
}

func replace(s string) string {
	return "bar"
}
