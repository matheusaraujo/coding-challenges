package main

func part2(puzzleInput []string) any {
	polymer := []byte(puzzleInput[0])
	minLen := len(polymer)

	for u := byte('a'); u <= 'z'; u++ {
		filtered := filterUnit(polymer, u)
		length := reactPolymer(filtered)

		if length < minLen {
			minLen = length
		}
	}

	return minLen
}

func filterUnit(polymer []byte, unit byte) []byte {
	result := make([]byte, 0, len(polymer))

	lower := unit
	upper := unit - 32

	for _, c := range polymer {
		if c != lower && c != upper {
			result = append(result, c)
		}
	}

	return result
}
