package main

func part2(puzzleInput []string) any {
	floor := 1
	for i, c := range puzzleInput[0] {
		if c == '(' {
			floor++
		}
		if c == ')' {
			floor--
		}

		if floor == -1 {
			return i
		}
	}
	return -1
}
