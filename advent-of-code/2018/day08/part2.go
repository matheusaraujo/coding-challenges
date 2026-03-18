package main

import (
	"strconv"
	"strings"
)

func part2(puzzleInput []string) any {
	fields := strings.Fields(puzzleInput[0])
	nums := make([]int, len(fields))
	for i, f := range fields {
		n, _ := strconv.Atoi(f)
		nums[i] = n
	}

	_, value := parseNode2(nums, 0)
	return value
}

func parseNode2(nums []int, i int) (int, int) {
	childCount := nums[i]
	metaCount := nums[i+1]
	i += 2

	childValues := make([]int, childCount)

	for c := 0; c < childCount; c++ {
		i, childValues[c] = parseNode2(nums, i)
	}

	value := 0

	if childCount == 0 {
		for m := 0; m < metaCount; m++ {
			value += nums[i]
			i++
		}
	} else {
		for m := 0; m < metaCount; m++ {
			idx := nums[i]
			i++
			if idx >= 1 && idx <= childCount {
				value += childValues[idx-1]
			}
		}
	}

	return i, value
}
