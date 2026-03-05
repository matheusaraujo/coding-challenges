package main

func part1(puzzleInput []string) any {
	result := 0

	for _, line := range puzzleInput {
		component := NewComponent(line)

		if component.Ng > component.Nr && component.Ng > component.Nb {
			result += component.Id
		}
	}

	return result
}
