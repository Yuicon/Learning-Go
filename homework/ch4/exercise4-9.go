// 报告输入文本中每个单词出现的频率
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	seen := make(map[string]int) // a set of strings
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		line := input.Text()
		seen[line]++
		fmt.Println(line)
	}
	fmt.Println(seen)
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}
