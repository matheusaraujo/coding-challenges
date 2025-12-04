package main

func part1(puzzleInput []string) interface{} {
	result := 0

	for i := 0; i < len(puzzleInput); i++ {
		for j := 0; j < len(puzzleInput[i]); j++ {
			if accessible(puzzleInput, i, j) {
				result++
			}
		}
	}

	return result
}
