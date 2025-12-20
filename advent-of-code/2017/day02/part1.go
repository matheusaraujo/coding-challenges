package main

import (
	"strconv"
)

func part1(puzzleInput []string) string {
	numbers := parseInput(puzzleInput)
	checksum := 0
	for _, row := range numbers {
		checksum += row[len(row)-1] - row[0]
	}
	return strconv.Itoa(checksum)
}
