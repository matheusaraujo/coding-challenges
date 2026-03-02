package main

func part2(puzzleInput []string) any {
	const size = 1000

	fabric := make([][]int, size)
	for i := range fabric {
		fabric[i] = make([]int, size)
	}

	claims := make([]claim, 0, len(puzzleInput))

	for _, line := range puzzleInput {
		c := parseLine(line)
		claims = append(claims, c)

		for y := c.top; y < c.top+c.height; y++ {
			for x := c.left; x < c.left+c.width; x++ {
				fabric[y][x]++
			}
		}
	}

	for _, c := range claims {
		overlaps := false

		for y := c.top; y < c.top+c.height && !overlaps; y++ {
			for x := c.left; x < c.left+c.width; x++ {
				if fabric[y][x] > 1 {
					overlaps = true
					break
				}
			}
		}

		if !overlaps {
			return c.id
		}
	}

	return nil
}
