package main

func part1(puzzleInput []string) any {
	points := parsePoints(puzzleInput)

	minX, maxX := points[0].x, points[0].x
	minY, maxY := points[0].y, points[0].y

	for _, p := range points {
		if p.x < minX {
			minX = p.x
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	areas := make([]int, len(points))
	infinite := make([]bool, len(points))

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {

			cell := Point{x, y}

			bestDist := int(^uint(0) >> 1)
			bestIdx := -1
			tie := false

			for i, p := range points {
				d := manhattan(cell, p)

				if d < bestDist {
					bestDist = d
					bestIdx = i
					tie = false
				} else if d == bestDist {
					tie = true
				}
			}

			if tie {
				continue
			}

			areas[bestIdx]++

			if x == minX || x == maxX || y == minY || y == maxY {
				infinite[bestIdx] = true
			}
		}
	}

	maxArea := 0
	for i := range areas {
		if !infinite[i] && areas[i] > maxArea {
			maxArea = areas[i]
		}
	}

	return maxArea
}