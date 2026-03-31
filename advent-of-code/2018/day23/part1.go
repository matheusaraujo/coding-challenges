package main

func part1(puzzleInput []string) any {
	bots := parseBots(puzzleInput)

	// Find strongest bot
	strongest := bots[0]
	for _, b := range bots {
		if b.r > strongest.r {
			strongest = b
		}
	}

	// Count bots in range
	count := 0
	for _, b := range bots {
		if manhattan(strongest.x, strongest.y, strongest.z, b.x, b.y, b.z) <= strongest.r {
			count++
		}
	}

	return count
}
