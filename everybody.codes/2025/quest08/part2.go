package main

import (
	"strconv"
)

func part2(puzzleInput []string) string {
	lines := parseInput(puzzleInput)
	intersections := 0

	for i := 1; i < len(lines); i++ {
		for j := 0; j < i; j++ {
			if intersects(lines[i].a, lines[i].b, lines[j].a, lines[j].b) {
				intersections++
			}
		}
	}

	return strconv.Itoa(intersections)
}
