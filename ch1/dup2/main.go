// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Exercise 1.4

func main() {
	counts := make(map[string]int)
	fnames := make(map[string]map[string]bool)
	files := os.Args[1:]

	for _, arg := range files {
		countLines(arg, counts, fnames)
	}

	for line, n := range counts {
		if n > 1 {
			fs := make([]string, 0, len(fnames[line]))
			for f, _ := range fnames[line] {
				fs = append(fs, f)
			}
			fmt.Printf("%d\t%s\t\t%s\n", n, line, strings.Join(fs, ","))
		}
	}
}

func countLines(file string, counts map[string]int, fnames map[string]map[string]bool) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
		return
	}
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if fnames[input.Text()] == nil {
			fnames[input.Text()] = make(map[string]bool)
		}
		fnames[input.Text()][file] = true
	}
	f.Close()
	// NOTE: ignoring potential errors from input.Err()
}

//!-
