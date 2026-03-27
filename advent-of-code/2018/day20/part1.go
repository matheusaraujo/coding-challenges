package main

func part1(puzzleInput []string) any {
	distances := buildMap(puzzleInput[0])
	maxDist := 0
	for _, d := range distances {
		if d > maxDist {
			maxDist = d
		}
	}
	return maxDist
}
