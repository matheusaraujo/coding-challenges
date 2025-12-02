package main

import (
	"strings"
)

func part2(puzzleInput []string) interface{} {
	return solve(puzzleInput, invalid2)
}

func invalid2(num string) bool {
	n := len(num)
	half := n / 2
	for i := half; i > 0; i -= 1 {
		s := num[0:i]
		a := strings.ReplaceAll(num, s, "")
		if a == "" {
			return true
		}
	}
	return false
}
