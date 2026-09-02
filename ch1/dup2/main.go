// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// Modified by Filipp Zapolskikh

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	file_lines := make(map[*os.File]map[string]bool)
	if len(files) == 0 {
		file_lines[os.Stdin] = make(map[string]bool)
		countLines(os.Stdin, counts, file_lines[os.Stdin])
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			file_lines[f] = make(map[string]bool)
			countLines(f, counts, file_lines[f])
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			var files_with_line []string
			for file := range file_lines {
				if file_lines[file][line] {
					files_with_line = append(files_with_line, file.Name())
				}
			}
			fmt.Printf("%d\t%v\t%s\n", n, files_with_line, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, lines map[string]bool) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		lines[input.Text()] = true
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
