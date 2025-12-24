package main

import (
	"strconv"
)

// pos_n = (pos_{n-1} + step) % n + 1
func part1(puzzleInput []string) any {
	step, _ := strconv.Atoi(puzzleInput[0])
	buf := []int{0}
	pos := 0

	for n := 1; n <= 2017; n++ {
		pos = (pos+step)%len(buf) + 1
		buf = append(buf[:pos], append([]int{n}, buf[pos:]...)...)
	}

	return buf[(pos+1)%len(buf)]
}
