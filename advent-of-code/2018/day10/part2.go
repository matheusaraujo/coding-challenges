package main

func part2(puzzleInput []string) any {
	points := parseInput(puzzleInput)

	bestTime := 0
	smallestHeight := 1 << 30

	// same simulation as part1
	for t := 0; ; t++ {
		ps := advance(points, t)
		_, _, minY, maxY := boundingBox(ps)
		height := maxY - minY
		if height < smallestHeight {
			smallestHeight = height
			bestTime = t
		} else {
			break
		}
	}

	return bestTime // seconds until message appears
}
