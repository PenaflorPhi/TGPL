// Experiment to measure the difference in running time between our potentially inefficient version and the one that sues strings.Join
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start_time := time.Now().Unix()
	fmt.Println(strings.Join(os.Args[1:], " "))
	final_time := time.Now().Unix()

	fmt.Println(final_time-start_time, "s")

	start_time = time.Now().Unix()

	var s, sep = "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	final_time = time.Now().Unix()

	fmt.Println(final_time-start_time, "s")

}
