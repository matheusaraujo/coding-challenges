package main

import (
	"strconv"
)

func part3(puzzleInput []string) string {
	return strconv.Itoa(solve(parseInput(puzzleInput)))
}
