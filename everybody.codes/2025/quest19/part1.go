package main

import (
	"strconv"
)

func part1(puzzleInput []string) string {
	return strconv.Itoa(solve(parseInput(puzzleInput)))
}
