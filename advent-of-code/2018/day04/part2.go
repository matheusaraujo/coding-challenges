package main

func part2(puzzleInput []string) any {
	sleepMap := parseInput(puzzleInput)

	bestGuard := 0
	bestMinute := 0
	maxCount := 0

	for guard, minutes := range sleepMap {
		for minute, count := range minutes {
			if count > maxCount {
				maxCount = count
				bestGuard = guard
				bestMinute = minute
			}
		}
	}

	return bestGuard * bestMinute
}
