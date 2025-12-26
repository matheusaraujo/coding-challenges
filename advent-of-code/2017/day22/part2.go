package main

func part2(puzzleInput []string) any {
	const (
		Clean = iota
		Weakened
		Infected
		Flagged
	)

	grid := make(map[[2]int]int)

	h := len(puzzleInput)
	w := len(puzzleInput[0])

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if puzzleInput[y][x] == '#' {
				grid[[2]int{x - w/2, y - h/2}] = Infected
			}
		}
	}

	dir := 0
	x, y := 0, 0
	infections := 0

	for i := 0; i < 10000000; i++ {
		pos := [2]int{x, y}
		state := grid[pos]

		switch state {
		case Clean:
			dir = (dir + 3) % 4
		case Weakened:

		case Infected:
			dir = (dir + 1) % 4
		case Flagged:
			dir = (dir + 2) % 4
		}

		switch state {
		case Clean:
			grid[pos] = Weakened
		case Weakened:
			grid[pos] = Infected
			infections++
		case Infected:
			grid[pos] = Flagged
		case Flagged:
			delete(grid, pos)
		}

		switch dir {
		case 0:
			y--
		case 1:
			x++
		case 2:
			y++
		case 3:
			x--
		}
	}

	return infections
}
