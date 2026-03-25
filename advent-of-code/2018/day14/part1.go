package main

import (
	"strconv"
)

func part1(puzzleInput []string) any {
	n, _ := strconv.Atoi(puzzleInput[0])

	recipes := []int{3, 7}
	e1, e2 := 0, 1

	for len(recipes) < n+10 {
		step(&recipes, &e1, &e2)
	}

	result := ""
	for i := n; i < n+10; i++ {
		result += strconv.Itoa(recipes[i])
	}

	return result
}
