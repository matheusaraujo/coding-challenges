package main

func part1(puzzleInput []string) any {
	barrels := parseInput(puzzleInput)
	return destroy(barrels, 0, 0)
}
