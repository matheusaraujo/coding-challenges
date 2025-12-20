package main

import (
	"strconv"
)

func part2(puzzleInput []string) string {
	points := parseInput(puzzleInput)
	n := len(points)

	pairs := buildPairs(points)
	d := newDSU(n)
	components := n

	for _, pr := range pairs {
		if d.union(pr.i, pr.j) {
			components--
			if components == 1 {
				return strconv.Itoa(points[pr.i].x * points[pr.j].x)
			}
		}
	}

	return "0"
}
