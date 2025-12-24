package main

func part1(puzzleInput []string) any {
	path := parseInput(puzzleInput)
	x, y, z := 0, 0, 0

	for _, step := range path {
		x, y, z = move(step, x, y, z)
	}

	return max(abs(x), max(abs(y), abs(z)))
}
