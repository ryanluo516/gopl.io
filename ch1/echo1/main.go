// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	// Exercise 1.1
	fmt.Println("command: " + os.Args[0])

	// Exercise 1.2
	fmt.Println("arguments: ")
	for i, arg := range os.Args[1:] {
		fmt.Printf("%d:\t%s\n", i, arg)
	}
}

//!-
