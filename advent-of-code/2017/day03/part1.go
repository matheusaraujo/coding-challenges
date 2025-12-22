package main

import (
	"math"
	"strconv"
)

func part1(puzzleInput []string) any {
	input, _ := strconv.Atoi(puzzleInput[0])
	return distance(input)
}

func distance(n int) int {
	k := int(math.Ceil((math.Sqrt(float64(n)) - 1) / 2))

	m := (2*k + 1) * (2*k + 1)
	d := m - n
	side := 2 * k

	var x, y int

	switch {
	case d < side:
		x = k - d
		y = -k
	case d < 2*side:
		x = -k
		y = -k + (d - side)
	case d < 3*side:
		x = -k + (d - 2*side)
		y = k
	default:
		x = k
		y = k - (d - 3*side)
	}

	return abs(x) + abs(y)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
