package main

import (
	"convert/converter"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "no arguments provided\n")
		os.Exit(1)
	}

	val, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	from := args[1]
	to := args[2]
	result, err := converter.Convert(val, from, to)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stdout, result)
}
