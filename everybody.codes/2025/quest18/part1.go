package main

import (
	"strconv"
)

func part1(puzzleInput []string) string {
	plants, id, _, _, _ := parseInput(puzzleInput)
	return strconv.Itoa(plants[id].Energy(id, plants))
}
