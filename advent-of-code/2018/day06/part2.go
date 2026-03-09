package main

func part2(puzzleInput []string) any {
	const limit = 10000

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

	margin := limit / len(points)

	minX -= margin
	maxX += margin
	minY -= margin
	maxY += margin

	regionSize := 0

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {

			cell := Point{x, y}

			totalDist := 0
			for _, p := range points {
				totalDist += manhattan(cell, p)
				if totalDist >= limit {
					break
				}
			}

			if totalDist < limit {
				regionSize++
			}
		}
	}

	return regionSize
}