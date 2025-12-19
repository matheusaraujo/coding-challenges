package main

import (
	"strconv"
)

func parseInptut(puzzleInput []string) []int {
	result := make([]int, len(puzzleInput))
	for i, line := range puzzleInput {
		n, _ := strconv.Atoi(line)
		result[i] = n
	}
	return result
}

func solve(puzzleInput []string, inc func(int) int) int {
	offsets := parseInptut(puzzleInput)
	jumps, idx := 0, 0
	for idx >= 0 && idx < len(offsets) {
		jumps++
		next := offsets[idx]
		offsets[idx] = inc(offsets[idx])
		idx += next
	}
	return jumps
}
