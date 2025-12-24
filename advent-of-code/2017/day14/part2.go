package main

func part2(puzzleInput []string) any {
	key := puzzleInput[0]
	grid := buildGrid(key)

	regions := 0
	for i := 0; i < 128; i++ {
		for j := 0; j < 128; j++ {
			if grid[i][j] {
				dfs(grid, i, j)
				regions++
			}
		}
	}

	return regions
}

func buildGrid(key string) [][]bool {
	grid := make([][]bool, 128)
	for i := 0; i < 128; i++ {
		row := getRow(key, i)
		hash := knotHash(row)
		binRow := hexToBin(hash)
		grid[i] = make([]bool, 128)
		for j, c := range binRow {
			if c == '1' {
				grid[i][j] = true
			}
		}
	}
	return grid
}

func dfs(grid [][]bool, x, y int) {
	if x < 0 || x >= 128 || y < 0 || y >= 128 {
		return
	}
	if !grid[x][y] {
		return
	}
	grid[x][y] = false // visited
	dfs(grid, x+1, y)
	dfs(grid, x-1, y)
	dfs(grid, x, y+1)
	dfs(grid, x, y-1)
}
