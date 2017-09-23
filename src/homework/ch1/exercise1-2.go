// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, value := range os.Args[1:] {
		fmt.Printf("%[1]s %[2]d \n", value, index)
	}
}
