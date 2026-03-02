package main

func part1(puzzleInput []string) any {
	const size = 1000

	fabric := make([][]int, size)
	for i := range fabric {
		fabric[i] = make([]int, size)
	}

	for _, line := range puzzleInput {
		c := parseLine(line)

		for y := c.top; y < c.top+c.height; y++ {
			for x := c.left; x < c.left+c.width; x++ {
				fabric[y][x]++
			}
		}
	}

	overlapCount := 0
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if fabric[y][x] >= 2 {
				overlapCount++
			}
		}
	}

	return overlapCount
}
