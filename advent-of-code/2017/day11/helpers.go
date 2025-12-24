package main

import (
	"strings"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func parseInput(puzzleInput []string) []string {
	return strings.Split(puzzleInput[0], ",")
}

func move(step string, x int, y int, z int) (int, int, int) {
	switch step {
	case "n":
		y--
		z++
	case "ne":
		x++
		y--
	case "se":
		x++
		z--
	case "s":
		y++
		z--
	case "sw":
		x--
		y++
	case "nw":
		x--
		z++
	}
	return x, y, z
}
