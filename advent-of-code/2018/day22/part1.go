package main

func part1(puzzleInput []string) any {
	depth, target := parseInput(puzzleInput)

	grid := buildGrid(depth, target, target.x, target.y)

	risk := 0
	for y := 0; y <= target.y; y++ {
		for x := 0; x <= target.x; x++ {
			risk += int(grid[y][x])
		}
	}

	return risk
}
