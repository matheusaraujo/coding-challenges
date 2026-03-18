package main

import (
	"fmt"
	"strconv"
)

// part2 finds the top-left and size of the square with largest total power efficiently
func part2(puzzleInput []string) any {
	serial, _ := strconv.Atoi(puzzleInput[0])
	_, sum := BuildGrid(serial)

	maxPower := -1 << 30
	bestX, bestY, bestSize := 0, 0, 0

	// iterate square sizes 1..300
	for size := 1; size <= 300; size++ {
		limit := 301 - size
		for y := 1; y <= limit; y++ {
			for x := 1; x <= limit; x++ {
				total := sum[y+size-1][x+size-1] - sum[y-1][x+size-1] - sum[y+size-1][x-1] + sum[y-1][x-1]
				if total > maxPower {
					maxPower = total
					bestX, bestY, bestSize = x, y, size
				}
			}
		}
		// optional: break early if maxPower not improving (tweakable)
	}

	return fmt.Sprintf("%d,%d,%d", bestX, bestY, bestSize)
}
