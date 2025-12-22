package main

func part2(puzzleInput []string) any {
	prev := -1
	curr := 0

	grid := make([]string, len(puzzleInput))
	copy(grid, puzzleInput)

	for prev != curr {
		prev = curr
		next := make([]string, len(grid))
		copy(next, grid)

		for i := 0; i < len(grid); i++ {
			row := []rune(grid[i])
			newRow := []rune(grid[i])

			for j := 0; j < len(row); j++ {
				if accessible(grid, i, j) {
					curr++
					newRow[j] = '.'
				} else {
					newRow[j] = row[j]
				}
			}
			next[i] = string(newRow)
		}
		grid = next
	}

	return curr
}
