package main

import (
	"strconv"
	"strings"
)

func part1(puzzleInput []string) string {
	openParentheses := strings.Count(puzzleInput[0], "(")
	closeParentheses := strings.Count(puzzleInput[0], ")")

	return strconv.Itoa(openParentheses - closeParentheses)
}
