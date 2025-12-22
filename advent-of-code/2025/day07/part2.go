package main

import (
	"strings"
)

func part2(puzzleInput []string) any {
	y := strings.Index(puzzleInput[0], "S")
	memo := make(map[[2]int]int)
	return down(puzzleInput, 1, y, memo)
}

func down(diagram []string, x, y int, memo map[[2]int]int) int {
	if y < 0 || y >= len(diagram[x]) {
		return 0
	}

	key := [2]int{x, y}
	if v, ok := memo[key]; ok {
		return v
	}

	if x == len(diagram)-1 {
		memo[key] = 1
		return 1
	}

	var res int
	if diagram[x][y] == '.' {
		res = down(diagram, x+1, y, memo)
	} else {
		res = down(diagram, x+1, y-1, memo) + down(diagram, x+1, y+1, memo)
	}

	memo[key] = res
	return res
}
