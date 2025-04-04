// Modify the `echo` program to print the index and value of each of its arguments, one per line.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep = "", ""
	for _, args := range os.Args[1:] {
		s += sep + args
		sep = " "
		fmt.Println(s)
	}
}
