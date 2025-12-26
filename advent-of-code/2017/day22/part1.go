package main

func part1(puzzleInput []string) any {
	infected := make(map[[2]int]bool)

	h := len(puzzleInput)
	w := len(puzzleInput[0])

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if puzzleInput[y][x] == '#' {
				infected[[2]int{x - w/2, y - h/2}] = true
			}
		}
	}

	dir := 0
	x, y := 0, 0
	infections := 0

	for i := 0; i < 10000; i++ {
		pos := [2]int{x, y}

		if infected[pos] {
			dir = (dir + 1) % 4
			delete(infected, pos)
		} else {
			dir = (dir + 3) % 4
			infected[pos] = true
			infections++
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
