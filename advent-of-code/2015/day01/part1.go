package main

import (
	"strings"
)

func part1(puzzleInput []string) interface{} {
	openParentheses := strings.Count(puzzleInput[0], "(")
	closeParentheses := strings.Count(puzzleInput[0], ")")

	return openParentheses - closeParentheses
}
