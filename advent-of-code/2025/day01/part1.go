package main

import (
	"strconv"
)

func part1(puzzleInput []string) string {
	dial, size, count := 50, 100, 0
	for _, rotation := range puzzleInput {
		dir, steps := parseInput(rotation)

		if dir == 'L' {
			steps = -steps
		}

		dial = mod(dial+steps, size)

		if dial == 0 {
			count++
		}
	}
	return strconv.Itoa(count)
}
