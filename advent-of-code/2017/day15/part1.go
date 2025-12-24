package main

func part1(puzzleInput []string) any {
	a, b := parseInput(puzzleInput)
	return solve(a, b, 40000000,
		func(x int) bool { return true },
		func(x int) bool { return true },
	)
}
