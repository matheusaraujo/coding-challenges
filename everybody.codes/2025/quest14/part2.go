package main

import (
	"strconv"
)

func part2(puzzleInput []string) string {
	floor := parseInput(puzzleInput)
	out := 0
	for i := 0; i < 2025; i++ {
		floor = next(floor)
		out += count(floor)
	}
	return strconv.Itoa(out)
}
