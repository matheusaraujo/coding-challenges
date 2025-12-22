package main

func part2(puzzleInput []string) any {
	_, loopSize := solve(puzzleInput)
	return loopSize
}
