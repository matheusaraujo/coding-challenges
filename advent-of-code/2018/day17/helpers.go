package main

import (
	"math"
	"regexp"
	"strconv"
)

type Point struct {
	x, y int
}

// solve runs the simulation and returns the answers for Part 1 and Part 2.
func solve(puzzleInput []string) (int, int) {
	grid := make(map[Point]rune)
	minY := math.MaxInt32
	maxY := math.MinInt32

	// Parse input strings like "x=495, y=2..7" or "y=7, x=495..501"
	re := regexp.MustCompile(`([xy])=(\d+), ([xy])=(\d+)\.\.(\d+)`)

	for _, line := range puzzleInput {
		match := re.FindStringSubmatch(line)
		if match == nil {
			continue
		}

		fixedVal, _ := strconv.Atoi(match[2])
		rangeStart, _ := strconv.Atoi(match[4])
		rangeEnd, _ := strconv.Atoi(match[5])

		if match[1] == "x" {
			for y := rangeStart; y <= rangeEnd; y++ {
				grid[Point{fixedVal, y}] = '#'
				if y < minY {
					minY = y
				}
				if y > maxY {
					maxY = y
				}
			}
		} else {
			y := fixedVal
			if y < minY {
				minY = y
			}
			if y > maxY {
				maxY = y
			}
			for x := rangeStart; x <= rangeEnd; x++ {
				grid[Point{x, y}] = '#'
			}
		}
	}

	if minY > maxY {
		return 0, 0
	}

	// Run the water simulation starting from the spring at (500, 0)
	fill(Point{500, 0}, grid, maxY)

	part1Count := 0
	part2Count := 0

	// Tally up the results within the strict minY to maxY boundary bounds
	for p, v := range grid {
		if p.y >= minY && p.y <= maxY {
			if v == '~' || v == '|' {
				part1Count++
			}
			if v == '~' {
				part2Count++
			}
		}
	}

	return part1Count, part2Count
}

// fill simulates water dropping downwards.
// It returns true if the water settles (is blocked by clay or settled water).
func fill(p Point, grid map[Point]rune, maxY int) bool {
	if p.y > maxY {
		return false // Drained into the abyss
	}

	val := grid[p]
	if val == '#' || val == '~' {
		return true // Hit a solid surface
	}
	if val == '|' {
		return false // Already flowing/draining here
	}

	// Mark as flowing water
	grid[p] = '|'

	down := Point{p.x, p.y + 1}
	belowBlocked := grid[down] == '#' || grid[down] == '~'
	if !belowBlocked {
		belowBlocked = fill(down, grid, maxY)
	}

	// If the block below is solid, try to spread left and right
	if belowBlocked {
		leftBlocked := spread(p, -1, grid, maxY)
		rightBlocked := spread(p, 1, grid, maxY)

		// If bounded by walls on both sides, this row settles
		if leftBlocked && rightBlocked {
			lx := p.x
			for grid[Point{lx, p.y}] == '|' {
				grid[Point{lx, p.y}] = '~'
				lx--
			}
			rx := p.x + 1
			for grid[Point{rx, p.y}] == '|' {
				grid[Point{rx, p.y}] = '~'
				rx++
			}
			return true
		}
	}
	return false
}

// spread recursively flows water horizontally.
func spread(p Point, dx int, grid map[Point]rune, maxY int) bool {
	curr := Point{p.x + dx, p.y}

	if grid[curr] == '#' {
		return true // Hit a clay wall
	}

	grid[curr] = '|'

	down := Point{curr.x, curr.y + 1}
	belowBlocked := grid[down] == '#' || grid[down] == '~'
	if !belowBlocked {
		belowBlocked = fill(down, grid, maxY)
	}

	// Keep spreading if the ground beneath is still solid
	if belowBlocked {
		return spread(curr, dx, grid, maxY)
	}

	return false
}
