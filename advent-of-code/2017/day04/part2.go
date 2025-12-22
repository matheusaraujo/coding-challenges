package main

import (
	"slices"
)

func part2(puzzleInput []string) any {
	return solve(puzzleInput, rearange)
}

func rearange(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	slices.Sort(runes)
	return string(runes)
}
