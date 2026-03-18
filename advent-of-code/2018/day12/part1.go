package main

func part1(puzzleInput []string) any {
	pots, rules := parseInputDay12(puzzleInput)

	generations := 20
	for i := 0; i < generations; i++ {
		pots = nextGeneration(pots, rules)
	}

	return sumPots(pots)
}
