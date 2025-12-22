package main

func part2(puzzleInput []string) any {
	dial, size, count := 50, 100, 0

	for _, line := range puzzleInput {
		dir, steps := parseInput(line)

		var distToZero int
		if dir == 'R' {
			distToZero = size - dial
		} else {
			distToZero = dial
		}

		if steps >= distToZero {
			count += (steps-distToZero)/size + 1
			if distToZero == 0 {
				count--
			}
		}

		if dir == 'L' {
			steps = -steps
		}

		dial = mod(dial+steps, size)
	}

	return count
}
