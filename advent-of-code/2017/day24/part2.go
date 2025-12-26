package main

func part2(puzzleInput []string) any {
	comps := parseInput(puzzleInput)
	used := make([]bool, len(comps))
	_, strength := backtracking2(0, used, comps)
	return strength
}

func backtracking2(currentPort int, used []bool, comps []component) (int, int) {
	maxLength := 0
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
			length, strength := backtracking2(nextPort, used, comps)
			length++
			strength += c.a + c.b

			if length > maxLength || (length == maxLength && strength > maxStrength) {
				maxLength = length
				maxStrength = strength
			}
			used[i] = false
		}
	}
	return maxLength, maxStrength
}
