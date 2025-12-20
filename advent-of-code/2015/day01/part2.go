package main

import (
	"strconv"
)

func part2(puzzleInput []string) string {
	floor := 0
	for i, c := range puzzleInput[0] {
		if c == '(' {
			floor++
		}
		if c == ')' {
			floor--
		}

		if floor == -1 {
			return strconv.Itoa(i + 1)
		}
	}
	return "0"
}
