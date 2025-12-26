package main

const (
	Clean = iota
	Weakened
	Infected
	Flagged
)

func solve(puzzleInput []string, states []int, turn map[int]int, bursts int) int {
	grid := make(map[[2]int]int)
	h := len(puzzleInput)
	w := len(puzzleInput[0])

	// Initialize infected nodes
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if puzzleInput[y][x] == '#' {
				grid[[2]int{x - w/2, y - h/2}] = Infected
			}
		}
	}

	x, y := 0, 0
	dir := 0
	infections := 0

	for i := 0; i < bursts; i++ {
		pos := [2]int{x, y}
		state := grid[pos] // defaults to 0 (Clean) if missing in map

		// Determine new direction
		dir = (dir + turn[state]) % 4

		// Update state
		newState := (state + 1) % len(states)
		grid[pos] = newState
		if newState == Infected {
			infections++
		}

		// Move forward
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
