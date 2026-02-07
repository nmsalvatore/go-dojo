package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
)

func argsToInts(args []string) ([]int, error) {
	var nums []int
	for _, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil {
			return nil, fmt.Errorf("converting arguments to integers: %w", err)
		}
		nums = append(nums, num)
	}
	return nums, nil
}

func run(w io.Writer, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("want at least 1 argument, got %d", len(args))
	}

	nums, err := argsToInts(args)
	if err != nil {
		return err
	}

	min := slices.Min(nums)
	max := slices.Max(nums)
	mean, _ := Mean(nums)
	median, _ := Median(nums)

	fmt.Fprintf(w, "Min: %d\nMax: %d\nMean: %.2f\nMedian: %.2f\n", min, max, mean, median)
	return nil
}

func main() {
	args := os.Args[1:]

	err := run(os.Stdout, args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
