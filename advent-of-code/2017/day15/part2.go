package main

func part2(puzzleInput []string) any {
	a, b := parseInput(puzzleInput)
	return solve(a, b, 5000000,
		func(x int) bool { return x%4 == 0 },
		func(x int) bool { return x%8 == 0 },
	)
}
