package main

import (
	"strconv"
	"strings"
)

func part1(puzzleInput []string) any {
	fields := strings.Fields(puzzleInput[0])
	nums := make([]int, len(fields))
	for i, f := range fields {
		n, _ := strconv.Atoi(f)
		nums[i] = n
	}

	_, sum := parseNode1(nums, 0)
	return sum
}

func parseNode1(nums []int, i int) (int, int) {
	childCount := nums[i]
	metaCount := nums[i+1]
	i += 2

	total := 0

	for c := 0; c < childCount; c++ {
		var childSum int
		i, childSum = parseNode1(nums, i)
		total += childSum
	}

	for m := 0; m < metaCount; m++ {
		total += nums[i]
		i++
	}

	return i, total
}
