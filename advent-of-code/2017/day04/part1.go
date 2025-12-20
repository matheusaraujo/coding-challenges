package main

import (
	"strconv"
)

func part1(puzzleInput []string) string {
	return strconv.Itoa(solve(puzzleInput, bypasse))
}

func bypasse(s string) string {
	return s
}
