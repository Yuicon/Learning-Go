// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"homework/ch4/github"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		sinceDay := time.Since(item.CreatedAt).Hours() / 24
		fmt.Println(sinceDay)
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
		if sinceDay < 30 {
			fmt.Println("under month\n ")
			continue
		}
		if sinceDay > 365 {
			fmt.Println("more year\n ")
			continue
		}
		fmt.Println("under year\n ")
	}
}
