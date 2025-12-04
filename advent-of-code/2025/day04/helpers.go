package main

func accessible(grid []string, i, j int) bool {

	if grid[i][j] != '@' {
		return false
	}

	dirs := [8][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	neighbors := 0
	for _, d := range dirs {
		ni, nj := i+d[0], j+d[1]
		if ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(grid[ni]) && grid[ni][nj] == '@' {
			neighbors++
		}
	}

	return neighbors < 4
}
