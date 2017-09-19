// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	counts := make(map[string]int)    // counts of Unicode characters
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			counts["Letter"]++
		}
		if unicode.IsMark(r) {
			counts["Mark"]++
		}
		if unicode.IsNumber(r) {
			counts["Number"]++
		}
		if unicode.IsPunct(r) {
			counts["Punct"]++
		}
		if unicode.IsSpace(r) {
			counts["Space"]++
		}
		if unicode.IsSymbol(r) {
			counts["Symbol"]++
		}
		if unicode.IsControl(r) {
			counts["Control"]++
		}
	}
	fmt.Println(counts)
	fmt.Printf("runeKind\tcount\n")
	for c, n := range counts {
		fmt.Printf("%s\t%d\n", c, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
