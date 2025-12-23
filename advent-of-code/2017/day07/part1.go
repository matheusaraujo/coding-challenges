package main

func part1(puzzleInput []string) any {
	computers := parseInput(puzzleInput)
	return findBottom(computers)
}
