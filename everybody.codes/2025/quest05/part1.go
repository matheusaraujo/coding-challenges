package main

import (
	"strconv"
)

func part1(puzzleInput []string) string {
	sword := buildSword(puzzleInput[0])
	return strconv.Itoa(sword.quality)
}
