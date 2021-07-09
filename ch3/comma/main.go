// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// Exercise 3.10
func comma1(s string) string {
	n := len(s) % 3

	var buf bytes.Buffer
	for i, c := range s {
		buf.WriteString(string(c))
		if (i+1-n)%3 == 0 && i+1 < len(s) {
			buf.WriteString(",")
		}
	}

	return buf.String()
}

// Exercise 3.11
func comma2(s string) string {
	var buf bytes.Buffer

	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		buf.WriteString(s[:1])
		s = s[1:]
	}

	dot := strings.LastIndex(s, ".")

	var suf string
	if dot > 0 {
		suf = s[dot:]
		s = s[:dot]
	}

	n := len(s) % 3

	for i, c := range s {
		buf.WriteString(string(c))
		if (i+1-n)%3 == 0 && i+1 < len(s) {
			buf.WriteString(",")
		}
	}

	buf.WriteString(suf)

	return buf.String()
}

// Exercise 3.12
func anagrams(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	for _, c := range str1 {
		if !strings.Contains(str2, string(c)) {
			return false
		}
	}

	return true
}

//!-
