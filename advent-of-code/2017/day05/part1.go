package main

func part1(puzzleInput []string) any {
	return solve(puzzleInput, func(n int) int { return n + 1 })
}
