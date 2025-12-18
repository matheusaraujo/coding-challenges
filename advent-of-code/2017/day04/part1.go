package main

func part1(puzzleInput []string) interface{} {
	return solve(puzzleInput, bypasse)
}

func bypasse(s string) string {
	return s
}
