package main

// brute force solution
func part2(puzzleInput []string) any {
	delay := 0

	for {
		caught := false
		for _, line := range puzzleInput {
			depth, rng := parse(line)
			cycle := 2 * (rng - 1)
			if (depth+delay)%cycle == 0 {
				caught = true
				break
			}
		}
		if !caught {
			return delay
		}
		delay++
	}
}
