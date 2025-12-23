package main

func part2(puzzleInput []string) any {
	computers := parseInput(puzzleInput)
	bottom := findBottom(computers)
	_, corrected := calcTotalWeight(bottom, computers)
	return corrected
}

func calcTotalWeight(name string, computers map[string]computer) (int, int) {
	comp := computers[name]

	if len(comp.holding) == 0 {
		return comp.weight, 0
	}

	childWeights := make([]int, len(comp.holding))

	for i, child := range comp.holding {
		w, corr := calcTotalWeight(child, computers)
		childWeights[i] = w
		if corr != 0 {
			return 0, corr
		}
	}

	weightCount := make(map[int]int)
	for _, w := range childWeights {
		weightCount[w]++
	}

	if len(weightCount) == 1 {
		total := comp.weight
		for _, w := range childWeights {
			total += w
		}
		return total, 0
	}

	var correctWeight, wrongWeight int
	var wrongIndex int

	for w, count := range weightCount {
		if count == 1 {
			wrongWeight = w
		} else {
			correctWeight = w
		}
	}

	for i, w := range childWeights {
		if w == wrongWeight {
			wrongIndex = i
			break
		}
	}

	wrongChild := computers[comp.holding[wrongIndex]]
	diff := correctWeight - wrongWeight
	correctedWeight := wrongChild.weight + diff

	return 0, correctedWeight
}
