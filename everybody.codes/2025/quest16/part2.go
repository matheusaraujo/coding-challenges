package main

import (
	"strconv"
)

func part2(puzzleInput []string) string {
	blocks := parseInput(puzzleInput)
	spell := findSpell(blocks)
	return strconv.Itoa(prod(spell))
}
