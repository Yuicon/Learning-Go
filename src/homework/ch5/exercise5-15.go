package main

import (
	"fmt"
)

func main() {
	i, s := max(1, 2, 3, 4)
	fmt.Printf("%d %s max\n", i, s)
	i, s = min(1, 2, 3, 4)
	fmt.Printf("%d %s min\n", i, s)
	i, s = max()
	fmt.Printf("%d %s max\n", i, s)
	i, s = min()
	fmt.Printf("%d %s min\n", i, s)
}

func max(vals ...int) (int, string) {
	max := 0
	err := ""
	if len(vals) == 0 {
		err = "请至少传入一个参数"
	} else {
		for _, val := range vals {
			if max < val {
				max = val
			}
		}
	}

	return max, err
}

func min(vals ...int) (int, string) {
	min := 0
	err := ""
	if len(vals) == 0 {
		err = "请至少传入一个参数"
	} else {
		for _, val := range vals {
			if min > val {
				min = val
			}
		}
	}
	return min, err
}
