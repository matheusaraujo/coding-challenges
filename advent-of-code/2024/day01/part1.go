package main

import (
	"math"
	"strconv"
)

func part1(puzzleInput []string) string {
	left, right := parseInput(puzzleInput)
	sum := 0
	for i := range left {
		sum += int(math.Abs(float64(left[i] - right[i])))
	}
	return strconv.Itoa(sum)
}
