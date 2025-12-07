package main

import (
	"strings"
)

func part1(puzzleInput []string) interface{} {
	y := strings.Index(puzzleInput[0], "S")
	return split(puzzleInput, y)
}

func split(diagram []string, y int) int {
	beams := make(map[int]bool)
	beams[y] = true
	split := 0

	for x := 1; x < len(diagram); x++ {
		next := make(map[int]bool, 0)
		for b, _ := range beams {
			if b < 0 || b >= len(diagram[x]) {
				continue
			}
			if diagram[x][b] == '.' {
				if _, ok := next[b]; !ok {
					next[b] = true
				}
			} else if diagram[x][b] == '^' {
				split++
				if _, ok := next[b-1]; !ok {
					next[b-1] = true
				}
				if _, ok := next[b+1]; !ok {
					next[b+1] = true
				}
			}
		}
		beams = next
	}

	return split
}
