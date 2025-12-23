package main

func part1(puzzleInput []string) any {
	lengths := parseCommaInts(puzzleInput[0])
	sparse := runKnot(lengths, 1)
	return sparse[0] * sparse[1]
}
