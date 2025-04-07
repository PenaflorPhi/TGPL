// Exercise 1.4: Modify dup2 to print the names of all files in which each duplicated line occurs
package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func countLines(f *os.File, lineCounts map[string]int, lineToFiles map[string][]string, fileName string) {
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		lineCounts[line]++

		if !slices.Contains(lineToFiles[line], fileName) {
			lineToFiles[line] = append(lineToFiles[line], fileName)
		}
	}
}

func printDups(lineCounts map[string]int, lineToFiles map[string][]string) {
	for line, count := range lineCounts {
		if count > 1 {
			fmt.Printf("Files: %v\n", lineToFiles[line])
			fmt.Printf("Count: %d\tLine: %s\n\n", count, line)
		}
	}
}

func main() {
	lineCounts := make(map[string]int)
	lineToFiles := make(map[string][]string)

	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, lineCounts, lineToFiles, "stdin")
	} else {
		for _, fileName := range files {
			f, err := os.Open(fileName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error opening file %s: %v\n", fileName, err)
				continue
			}
			countLines(f, lineCounts, lineToFiles, fileName)
			f.Close()
		}
	}

	printDups(lineCounts, lineToFiles)
}
