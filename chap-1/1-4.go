// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
		printDups(counts)
	} else {
		for _, arg := range files {
			counts = make(map[string]int)
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			countLines(f, counts)
			f.Close()

			has_duplicates := hasDups(counts)

			if has_duplicates {
				fmt.Println("file: " + arg)
			} else {
				fmt.Println("file: " + arg + " has no duplicates!")
			}
			printDups(counts)
		}
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

func printDups(counts map[string]int) {
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%v\n", n, line)
		}
	}
}

func hasDups(counts map[string]int) bool {
	dup_flag := false
	for _, n := range counts {
		if n > 1 {
			dup_flag = true
		}
	}
	return dup_flag
}
