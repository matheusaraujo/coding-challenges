package main

import (
	"strconv"
)

func part1(puzzleInput []string) string {
	floor := parseInput(puzzleInput)
	out := 0
	for i := 0; i < 10; i++ {
		floor = next(floor)
		out += count(floor)
	}
	return strconv.Itoa(out)
}
