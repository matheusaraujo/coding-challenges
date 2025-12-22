package main

import (
	"math"
)

func part1(puzzleInput []string) any {
	left, right := parseInput(puzzleInput)
	sum := 0
	for i := range left {
		sum += int(math.Abs(float64(left[i] - right[i])))
	}
	return sum
}
