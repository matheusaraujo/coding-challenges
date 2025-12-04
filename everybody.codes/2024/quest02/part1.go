package main

import (
	"strings"
)

func part1(puzzleInput []string) interface{} {
	runics, inscriptions := parseInput(puzzleInput)
	result := 0

	for _, r := range runics {
		result += strings.Count(inscriptions[0], r)
	}
	return result
}
