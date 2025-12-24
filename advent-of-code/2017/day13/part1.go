package main

func part1(puzzleInput []string) any {
	severity := 0

	for _, line := range puzzleInput {
		depth, rng := parse(line)
		cycle := 2 * (rng - 1)
		if depth%cycle == 0 {
			severity += depth * rng
		}
	}

	return severity
}
