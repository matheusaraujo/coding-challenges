package main

func part2(puzzleInput []string) any {
	distances := buildMap(puzzleInput[0])
	count := 0
	for _, d := range distances {
		if d >= 1000 {
			count++
		}
	}
	return count
}
