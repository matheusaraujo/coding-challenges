package main

func part1(puzzleInput []string) any {
	numbers := parseInput(puzzleInput)
	checksum := 0
	for _, row := range numbers {
		checksum += row[len(row)-1] - row[0]
	}
	return checksum
}
