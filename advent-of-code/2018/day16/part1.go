package main

func part1(puzzleInput []string) any {
	samples, _ := parseInput(puzzleInput)
	threeOrMore := 0

	for _, s := range samples {
		matches := 0
		for _, opFunc := range Opcodes {
			if opFunc(s.Before, s.Inst) == s.After {
				matches++
			}
		}

		if matches >= 3 {
			threeOrMore++
		}
	}

	return threeOrMore
}
