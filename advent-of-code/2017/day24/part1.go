package main

func part1(puzzleInput []string) any {
	comps := parseInput(puzzleInput)
	used := make([]bool, len(comps))
	return backtracking(0, used, comps)
}

func backtracking(currentPort int, used []bool, comps []component) int {
	maxStrength := 0
	for i, c := range comps {
		if used[i] {
			continue
		}
		if c.a == currentPort || c.b == currentPort {
			used[i] = true
			nextPort := c.a
			if c.a == currentPort {
				nextPort = c.b
			} else {
				nextPort = c.a
			}
			strength := c.a + c.b + backtracking(nextPort, used, comps)
			if strength > maxStrength {
				maxStrength = strength
			}
			used[i] = false
		}
	}
	return maxStrength
}
