package main

import (
	"strconv"
)

func parseInput(line string) (byte, int) {
	dir, ns := line[0], line[1:]
	steps, _ := strconv.Atoi(ns)

	return dir, steps
}

// go % funcion returns negative when x is negative,
// this function returns positive
func mod(x, y int) int {
	a := x % y
	if a < 0 {
		a += y
	}
	return a
}
