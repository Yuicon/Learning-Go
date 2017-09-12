package main

import (
	"fmt"
	"bufio"
	"os"
	"crypto/sha256"
	"crypto/sha512"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if os.Args[1] == "Sum384" {
			fmt.Printf("input Sum384 %s %x\n", input.Text(), sha512.Sum384([]byte(input.Text())))
			continue
		}
		if os.Args[1] == "Sum512" {
			fmt.Printf("input Sum512 %s %x\n", input.Text(), sha512.Sum384([]byte(input.Text())))
			continue
		}
		fmt.Printf("input Sum256 %s %x\n", input.Text(), sha256.Sum256([]byte(input.Text())))
	}
}
