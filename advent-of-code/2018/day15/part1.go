package main

func part1(puzzleInput []string) any {
	_, score := solve(puzzleInput, 3, false)
	return score
}
