package main

func part2(puzzleInput []string) any {
	result := NewComponent(puzzleInput[0])

	for _, line := range puzzleInput {
		component := NewComponent(line)

		if component.Ns > result.Ns ||
			(component.Ns == result.Ns && component.Nc < result.Nc) {
			result = component
		}
	}

	return result.Id
}
