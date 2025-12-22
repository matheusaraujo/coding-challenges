package main

// dumb brute force solution
func part2(puzzleInput []string) any {
	n := len(puzzleInput)
	xc, yc, rmx, tmx := n/2, n/2, 0, 0

	prevTot := 0
	for r := 1; r <= n/2; r++ {
		tot := 0
		for x := 0; x < n; x++ {
			for y := 0; y < n; y++ {
				if innerCircle(x, y, xc, yc, r) {
					tot += atoi(puzzleInput, x, y)
				}
			}
		}
		tot -= prevTot
		if tot > tmx {
			tmx = tot
			rmx = r
		}
		prevTot += tot
	}

	return tmx * rmx
}
