package main

import (
	"strconv"
)

func part2(puzzleInput []string) string {
	barrels := parseInput(puzzleInput)
	return strconv.Itoa(destroy(barrels, 0, 0) + destroy(barrels, len(barrels)-1, len(barrels[0])-1))
}
