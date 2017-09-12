package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%t\n", isReorder("abcdefg", "gfedcba"))
}

func isReorder(l string, r string) bool {
	if (l == r || len(l) != len(r)) {
		fmt.Printf("%s %s\n", l, r)
		return false
	}
	runeArray := makeSortRuneArray(l)
	fmt.Printf("%s  %q\n", l, runeArray)
	runeArray2 := makeSortRuneArray(r)
	fmt.Printf("%s  %q\n", r, runeArray2)
	return equal(runeArray, runeArray2)
}

func makeSortRuneArray(s string) []rune {
	runeArray := make([]rune, len(s))
	for i, t := range s {
		runeArray[i] = t
	}
	quickSort(runeArray, 0, len(s) - 1)
	return runeArray
}

func equal(x, y []rune) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func swap(a, b *rune) {
	tmp := *a;
	*a = *b;
	*b = tmp;
}

func partition(runeArray []rune, low, high int) int {
	privotKey := runeArray[low]
	for low < high {
		for low < high  && runeArray[high] >= privotKey {
			high--
		}
		swap(&runeArray[low], &runeArray[high])
		for low < high  && runeArray[low] <= privotKey {
			low++
		}
		swap(&runeArray[low], &runeArray[high])
	}
	return low
}

func quickSort(runeArray []rune, low, high int) {
	if(low < high){
		privotLoc := partition(runeArray,  low,  high);  //将表一分为二
		quickSort(runeArray,  low,  privotLoc -1);          //递归对低子表递归排序
		quickSort(runeArray,   privotLoc + 1, high);        //递归对高子表递归排序
	}
}