package main

func part2(puzzleInput []string) any {
	path := parseInput(puzzleInput)
	x, y, z, m := 0, 0, 0, 0

	for _, step := range path {
		x, y, z = move(step, x, y, z)
		m = max(m, abs(x), max(abs(y), abs(z)))
	}

	return m
}
