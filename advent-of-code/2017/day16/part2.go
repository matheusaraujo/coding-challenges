package main

// cycle detection
func part2(puzzleInput []string) any {
	moves, dancers := parseInput(puzzleInput)

	seen := make(map[string]int)
	states := []string{}

	for i := 0; ; i++ {
		state := string(dancers)

		if firstSeen, ok := seen[state]; ok {
			cycleStart := firstSeen
			cycleLen := i - firstSeen

			remaining := (1_000_000_000 - cycleStart) % cycleLen
			return states[cycleStart+remaining]
		}

		seen[state] = i
		states = append(states, state)
		dancers = dance(dancers, moves)
	}
}
