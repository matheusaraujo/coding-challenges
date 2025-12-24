package main

func part2(puzzleInput []string) any {
	_, steps := path(puzzleInput)
	return steps
}
