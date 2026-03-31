package main

func part1(puzzleInput []string) any {
	_, units, _ := simulate(puzzleInput, 0)
	return units
}
