package main

import (
	"strconv"
)

func part1(puzzleInput []string) string {
	return strconv.Itoa(solve(puzzleInput, func(n int) int { return n + 1 }))
}
