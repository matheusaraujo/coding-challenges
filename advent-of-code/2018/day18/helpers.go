package main

type Grid [][]byte

// solve runs the lumber collection simulation.
func solve(puzzleInput []string) (int, int) {
	grid := make(Grid, len(puzzleInput))
	for i, line := range puzzleInput {
		grid[i] = []byte(line)
	}

	seen := make(map[string]int)
	seen[gridString(grid)] = 0

	p1 := 0
	p2 := 0

	for m := 1; m <= 1000000000; m++ {
		grid = step(grid)

		if m == 10 {
			p1 = score(grid)
		}

		// Only look for cycles after Part 1 is safely computed
		if m > 10 {
			s := gridString(grid)
			if prevM, ok := seen[s]; ok {
				cycleLen := m - prevM
				remaining := 1000000000 - m
				fastForward := remaining % cycleLen

				for i := 0; i < fastForward; i++ {
					grid = step(grid)
				}
				p2 = score(grid)
				break
			}
			seen[s] = m
		}
	}

	return p1, p2
}

// step simulates a single minute of the lumber collection area.
func step(g Grid) Grid {
	R := len(g)
	C := len(g[0])
	next := make(Grid, R)

	for i := 0; i < R; i++ {
		next[i] = make([]byte, C)
		for j := 0; j < C; j++ {
			trees, lumber := 0, 0

			// Count 8 adjacent neighbors
			for dr := -1; dr <= 1; dr++ {
				for dc := -1; dc <= 1; dc++ {
					if dr == 0 && dc == 0 {
						continue
					}
					r, c := i+dr, j+dc
					if r >= 0 && r < R && c >= 0 && c < C {
						if g[r][c] == '|' {
							trees++
						}
						if g[r][c] == '#' {
							lumber++
						}
					}
				}
			}

			// Apply rules
			next[i][j] = g[i][j]
			if g[i][j] == '.' && trees >= 3 {
				next[i][j] = '|'
			} else if g[i][j] == '|' && lumber >= 3 {
				next[i][j] = '#'
			} else if g[i][j] == '#' {
				if lumber >= 1 && trees >= 1 {
					next[i][j] = '#'
				} else {
					next[i][j] = '.'
				}
			}
		}
	}
	return next
}

// gridString flattens the 2D grid into a string to be used as a map key for cycle detection.
func gridString(g Grid) string {
	var b []byte
	for _, row := range g {
		b = append(b, row...)
	}
	return string(b)
}

// score calculates the final resource value: total trees * total lumberyards.
func score(g Grid) int {
	trees, lumber := 0, 0
	for _, row := range g {
		for _, cell := range row {
			if cell == '|' {
				trees++
			} else if cell == '#' {
				lumber++
			}
		}
	}
	return trees * lumber
}
