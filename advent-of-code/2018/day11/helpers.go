package main

// ComputePower calculates the power level of a single cell
func ComputePower(x, y, serial int) int {
	rackID := x + 10
	power := rackID * y
	power += serial
	power *= rackID
	// extract hundreds digit
	power = (power / 100) % 10
	power -= 5
	return power
}

// BuildGrid builds a 2D grid of power levels and its summed-area table
func BuildGrid(serial int) (grid [301][301]int, sumTable [301][301]int) {
	for y := 1; y <= 300; y++ {
		for x := 1; x <= 300; x++ {
			grid[y][x] = ComputePower(x, y, serial)
			// summed-area table
			sumTable[y][x] = grid[y][x] + sumTable[y-1][x] + sumTable[y][x-1] - sumTable[y-1][x-1]
		}
	}
	return
}

// SquareSum returns the total power of a square from (x,y) with size s using summed-area table
func SquareSum(sumTable [301][301]int, x, y, size int) int {
	x2 := x + size - 1
	y2 := y + size - 1
	return sumTable[y2][x2] - sumTable[y-1][x2] - sumTable[y2][x-1] + sumTable[y-1][x-1]
}
