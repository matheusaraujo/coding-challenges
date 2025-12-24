package main

func part1(puzzleInput []string) any {
	moves, dancers := parseInput(puzzleInput)
	return string(dance(dancers, moves))
}
