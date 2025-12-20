package main

import (
	"strconv"
)

func part1(puzzleInput []string) string {
	n := len(puzzleInput)
	xc, yc, res := n/2, n/2, 0

	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			if innerCircle(x, y, xc, yc, 10) {
				res += atoi(puzzleInput, x, y)
			}
		}
	}

	return strconv.Itoa(res)
}
