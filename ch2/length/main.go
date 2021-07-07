// Exercise 2.2
package main

import (
	"bufio"
	"fmt"
	"gopl.io/ch2/lengthconv"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		for _, arg := range args {
			convertLength(arg)
		}
	} else {
		in := bufio.NewScanner(os.Stdin)
		for in.Scan() {
			convertLength(in.Text())
		}
	}
}

func convertLength(str string) {
	n, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "length: %v\n", err)
		os.Exit(1)
	}

	mn, fn := lengthconv.Meters(n), lengthconv.Feet(n)
	fmt.Fprintf(os.Stdout, "%s = %s, %s = %s\n", mn, lengthconv.MToF(mn), fn, lengthconv.FToM(fn))
}
