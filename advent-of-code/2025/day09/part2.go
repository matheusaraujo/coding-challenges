package main

func part2(puzzleInput []string) interface{} {
	points := parseInput(puzzleInput)
	mx, n := 0, len(points)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			p1, p2 := points[i], points[j]
			a := area(p1, p2)

			if a > mx && !intersects(p1, p2, points) {
				mx = max(a, mx)
			}
		}
	}

	return mx
}

func intersects(p1, p2 Point, points []Point) bool {
	n := len(points)
	for k := 0; k < n; k++ {
		l1, l2 := points[k], points[(k+1)%n]
		if intersectsLine(p1, p2, l1, l2) {
			return true
		}
	}
	return false
}

func intersectsLine(p1, p2, l1, l2 Point) bool {
	p_x_min, p_x_max := minMax(p1.x, p2.x)
	p_y_min, p_y_max := minMax(p1.y, p2.y)
	l_x_min, l_x_max := minMax(l1.x, l2.x)
	l_y_min, l_y_max := minMax(l1.y, l2.y)

	return l_x_max > p_x_min && l_x_min < p_x_max && l_y_max > p_y_min && l_y_min < p_y_max
}

func minMax(a, b int) (int, int) {
	return min(a, b), max(a, b)
}
