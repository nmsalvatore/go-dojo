package main

import (
	"errors"
	"sort"
)

func Mean(nums []int) (float64, error) {
	if len(nums) == 0 {
		return 0, errors.New("no arguments provided")
	}

	return (float64(Sum(nums)) / float64(len(nums))), nil
}

func Median(nums []int) (float64, error) {
	if len(nums) == 0 {
		return 0, errors.New("no arguments provided")
	}

	sorted := make([]int, len(nums))
	copy(sorted, nums)

	sort.Ints(sorted)
	count := len(sorted)

	if count%2 != 0 {
		i := count / 2
		return float64(sorted[i]), nil
	}

	v1 := sorted[(count/2)-1]
	v2 := sorted[count/2]
	mean, _ := Mean([]int{v1, v2})
	return mean, nil
}

func Sum(nums []int) (sum int) {
	for _, num := range nums {
		sum += num
	}
	return
}
