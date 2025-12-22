package main

func part1(puzzleInput []string) any {
	return solve(puzzleInput, bypasse)
}

func bypasse(s string) string {
	return s
}
