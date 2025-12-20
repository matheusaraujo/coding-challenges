package main

import (
	"strconv"
)

func part1(puzzleInput []string) string {
	barrels := parseInput(puzzleInput)
	return strconv.Itoa(destroy(barrels, 0, 0))
}
