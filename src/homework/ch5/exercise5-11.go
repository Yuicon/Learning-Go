package main

import (
	"fmt"
)

// prereqs记录了每个课程的前置课程
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"linear algebra":   {"calculus"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort2(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort2(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var isRing = func(items []string, key, value string) {
		for _, item := range items {
			if item == key {
				fmt.Printf("%s, %s 成环\n", key, value)
			}
		}
	}
	var visitAll func(items []string, key string)
	visitAll = func(items []string, key string) {
		for _, item := range items {
			// 如果没有被查看过
			if !seen[item] {
				//设置为已查看
				seen[item] = true
				//如果有前置课程就继续
				isRing(m[item], key, item)
				visitAll(m[item], item)
				//没有前置课程就加入学习列表
				order = append(order, item)
			}
		}
	}
	for key := range m {
		//随机从一门课的前置课程数组开始
		visitAll(m[key], key)
		//将不是任何课程前置的课程加入学习列表
		if !seen[key] {
			seen[key] = true
			order = append(order, key)
		}
	}
	return order
}
