package main

import (
	"strconv"
)

func part1(puzzleInput []string) string {
	output, columns := 0, 90
	blocks := parseInput(puzzleInput)
	for _, b := range blocks {
		output += columns / b
	}
	return strconv.Itoa(output)
}
