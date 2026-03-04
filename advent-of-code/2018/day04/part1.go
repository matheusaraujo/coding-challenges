package main

func part1(puzzleInput []string) any {
	sleepMap := parseInput(puzzleInput)

	sleepiestGuard := 0
	maxTotal := 0

	for guard, minutes := range sleepMap {
		total := 0
		for _, count := range minutes {
			total += count
		}
		if total > maxTotal {
			maxTotal = total
			sleepiestGuard = guard
		}
	}

	bestMinute := 0
	maxCount := 0
	for m, count := range sleepMap[sleepiestGuard] {
		if count > maxCount {
			maxCount = count
			bestMinute = m
		}
	}

	return sleepiestGuard * bestMinute
}
