package main

import (
	"strings"
)

func parseInput(puzzleInput []string) ([]string, []string) {
	runics := strings.Split(puzzleInput[0][6:], ",")
	return runics, puzzleInput[2:]
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
