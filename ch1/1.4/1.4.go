// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.

// Modify dupe2 to print the names of all files in which each duplicated name occurs

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files_counts := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, files_counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, files_counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, files_counts[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, file_counts map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		file_counts[input.Text()] = file_counts[input.Text()] + " " + f.Name()
	}
	// NOTE: ignoring potential errors from input.Err()
}
