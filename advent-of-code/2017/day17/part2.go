package main

import (
	"strconv"
)

func part2(puzzleInput []string) any {
	step, _ := strconv.Atoi(puzzleInput[0])
	pos, result := 0, 0

	for n := 1; n <= 50_000_000; n++ {
		pos = (pos+step)%n + 1
		if pos == 1 {
			result = n
		}
	}

	return result
}
