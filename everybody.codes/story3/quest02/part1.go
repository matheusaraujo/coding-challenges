package main

func part1(puzzleInput []string) any {

	start, bones := ParseGrid(puzzleInput)

	var bone Pos
	for b := range bones {
		bone = b
		break
	}

	visited := map[Pos]bool{}
	pos := start
	visited[pos] = true

	dirs := []Pos{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	dir := 0
	steps := 0

	for pos != bone {

		for {
			next := Add(pos, dirs[dir])

			if !visited[next] {

				pos = next
				visited[pos] = true
				steps++
				dir = (dir + 1) % 4
				break
			}

			dir = (dir + 1) % 4
		}
	}

	return steps
}
