package main

import (
	"fmt"
	"unicode/utf8"
	"bytes"
)

func main() {
	a := []byte("我拍了")
	a = reverseByte(a)
	fmt.Printf("%s \n", a)
}

func reverseByte(s []byte) []byte {
	var buf bytes.Buffer
	for j := len(s); j > 0; {
		r, size := utf8.DecodeLastRune(s[:j])
		buf.WriteRune(r)
		fmt.Printf("%d\t%c\n", j, r)
		j -= size
	}
	return buf.Bytes()
}
