package main

import (
	"fmt"
	"unicode"
	"bytes"
)

func main() {
	a := []byte("m   a x   ")
	a = replaceSpace(a)
	fmt.Printf("%s \n", a)
}

func replaceSpace(s []byte) []byte {
	var tmp rune
	var buf bytes.Buffer
	buf.Write(s)
	st := buf.String()
	buf.Reset()
	for _, t := range st {
		if !unicode.IsSpace(tmp) || !unicode.IsSpace(t) {
			buf.WriteRune(t)
		}
		tmp = t
	}
	return buf.Bytes()
}
