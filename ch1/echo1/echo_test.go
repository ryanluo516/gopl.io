package main_test

import (
	"fmt"
	"strings"
	"testing"
)

// Exercise 1.3 benchmark test

func echo2(args []string) {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3(args []string) {
	fmt.Println(strings.Join(args, " "))
}

var args = []string{"arg1", "arg2", "arg3", "arg4", "arg5", "arg6", "arg7", "arg8"}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2(args)
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo3(args)
	}
}
