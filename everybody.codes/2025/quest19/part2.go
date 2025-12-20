package main

import (
	"strconv"
)

func part2(puzzleInput []string) string {
	return strconv.Itoa(solve(parseInput(puzzleInput)))
}
