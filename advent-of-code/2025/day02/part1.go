package main

func part1(puzzleInput []string) any {
	return solve(puzzleInput, invalid1)
}

func invalid1(num string) bool {
	n := len(num)
	if n%2 != 0 {
		return false
	}
	half := n / 2
	return num[:half] == num[half:]
}
