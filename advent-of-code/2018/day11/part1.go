package main

import (
	"fmt"
	"strconv"
)

// part1 finds the top-left of the 3x3 square with largest total power
func part1(puzzleInput []string) any {
	serial, _ := strconv.Atoi(puzzleInput[0])
	_, sum := BuildGrid(serial)

	maxPower := -1 << 30
	bestX, bestY := 0, 0

	for y := 1; y <= 298; y++ { // 300-3+1 = 298
		for x := 1; x <= 298; x++ {
			// inline SquareSum for speed
			total := sum[y+2][x+2] - sum[y-1][x+2] - sum[y+2][x-1] + sum[y-1][x-1]
			if total > maxPower {
				maxPower = total
				bestX, bestY = x, y
			}
		}
	}

	return fmt.Sprintf("%d,%d", bestX, bestY)
}
