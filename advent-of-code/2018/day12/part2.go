package main

func part2(puzzleInput []string) any {
	pots, rules := parseInputDay12(puzzleInput)

	var lastSum, lastDiff int
	stableCount := 0
	generations := 50000000000 // 50 billion

	for gen := 1; gen <= generations; gen++ {
		pots = nextGeneration(pots, rules)
		sum := sumPots(pots)
		diff := sum - lastSum

		if diff == lastDiff {
			stableCount++
			if stableCount > 5 {
				// pattern is stable, extrapolate
				remaining := generations - gen
				return sum + remaining*diff
			}
		} else {
			stableCount = 0
		}

		lastSum = sum
		lastDiff = diff
	}

	return sumPots(pots)
}
