package main

import (
	"strconv"
	"strings"
)

func part1(puzzleInput []string) string {
	runics, inscriptions := parseInput(puzzleInput)
	result := 0

	for _, r := range runics {
		result += strings.Count(inscriptions[0], r)
	}
	return strconv.Itoa(result)
}
