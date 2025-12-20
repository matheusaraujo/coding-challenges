package main

import (
	"strconv"
)

func part2(puzzleInput []string) string {
	names, m := parseInput(puzzleInput)
	result := 0
	for i, name := range names {
		if isValid(name, m) {
			result += (i + 1)
		}
	}

	return strconv.Itoa(result)
}
