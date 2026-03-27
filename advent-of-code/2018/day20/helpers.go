package main

type Point struct {
	x, y int
}

// buildMap parses the regex and returns a map of Room -> Distance from start
func buildMap(regex string) map[Point]int {
	distances := make(map[Point]int)
	stack := []Point{}
	curr := Point{0, 0}
	distances[curr] = 0

	for _, char := range regex {
		switch char {
		case '^', '$':
			continue
		case '(':
			// Start of a branch: save current position
			stack = append(stack, curr)
		case '|':
			// New option in branch: go back to the start of this branch
			curr = stack[len(stack)-1]
		case ')':
			// End of branch: pop the stack
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		default:
			// Moving N, S, E, or W
			prev := curr
			switch char {
			case 'N':
				curr.y--
			case 'S':
				curr.y++
			case 'E':
				curr.x++
			case 'W':
				curr.x--
			}

			// If this room is new or we found a shorter way to it
			dist := distances[prev] + 1
			if d, exists := distances[curr]; !exists || dist < d {
				distances[curr] = dist
			}
		}
	}
	return distances
}
