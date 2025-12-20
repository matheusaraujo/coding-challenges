package main

import (
	"strconv"
	"strings"
)

func part1(puzzleInput []string) string {
	numbers, operations := parseInput(puzzleInput)
	result := 0
	for i, op := range operations {
		column := numbers[0][i]
		for j := 1; j < len(numbers); j++ {
			if op == "*" {
				column *= numbers[j][i]
			} else {
				column += numbers[j][i]
			}
		}
		result += column
	}
	return strconv.Itoa(result)
}

func parseInput(puzzleInput []string) ([][]int, []string) {
	numbers := make([][]int, len(puzzleInput)-1)
	for i := 0; i < len(puzzleInput)-1; i++ {
		line := strings.Split(puzzleInput[i], " ")
		numbers[i] = make([]int, 0)
		for _, n := range line {
			if len(strings.TrimSpace(n)) != 0 {
				m, _ := strconv.Atoi(n)
				numbers[i] = append(numbers[i], m)
			}
		}
	}
	line := strings.ReplaceAll(puzzleInput[len(puzzleInput)-1], " ", "")
	operations := strings.Split(line, "")

	return numbers, operations
}
