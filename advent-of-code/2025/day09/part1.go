package main

func part1(puzzleInput []string) any {
	mx := 0
	points := parseInput(puzzleInput)

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			mx = max(mx, area(points[i], points[j]))
		}
	}

	return mx
}
