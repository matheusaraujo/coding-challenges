package main

func part2(puzzleInput []string) any {
	return solve(puzzleInput, func(n int) int {
		if n >= 3 {
			return n - 1
		}
		return n + 1
	})
}
