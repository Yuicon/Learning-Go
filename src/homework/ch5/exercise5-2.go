//记录在HTML树中出现的同名元素的次数。
package main

import (
	"fmt"
	"os"
	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	//
	//for _, nodes := range visit2(make(map[string]int), doc) {
	//	fmt.Println(nodes)
	//}
	nodes := visit2(make(map[string]int), doc)
	fmt.Println(nodes)
}

// visit appends to links each link found in n and returns the result.
func visit2(nodes map[string]int, n *html.Node) map[string]int {
	//if n.Type == html.ElementNode && n.Data == "a" {
	//	for _, a := range n.Attr {
	//		if a.Key == "href" {
	//			nodes = append(nodes, a.Val)
	//		}
	//	}
	//}
	if n.Type == html.ElementNode {
		nodes["ElementNode"]++
	}
	if n.Type == html.TextNode {
		nodes["TextNode"]++
	}
	if n.Type == html.DocumentNode {
		nodes["DocumentNode"]++
	}
	if n.Type == html.ErrorNode {
		nodes["ErrorNode"]++
	}
	if n.Type == html.CommentNode {
		nodes["CommentNode"]++
	}
	if n.Type == html.DoctypeNode {
		nodes["DoctypeNode"]++
	}
	c := n.FirstChild
	if c != nil {
		c = c.NextSibling
		return visit2(nodes, c)
	}
	return nodes
}