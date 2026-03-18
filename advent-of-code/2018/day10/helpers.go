package main

import (
	"regexp"
	"strconv"
)

// Point represents a star with position and velocity
type Point struct {
	x, y   int
	vx, vy int
}

// parseInput parses input lines like:
// position=< 9,  1> velocity=< 0,  2>
func parseInput(lines []string) []Point {
	re := regexp.MustCompile(`position=<\s*(-?\d+),\s*(-?\d+)> velocity=<\s*(-?\d+),\s*(-?\d+)>`)
	points := make([]Point, 0, len(lines))

	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if matches != nil {
			x, _ := strconv.Atoi(matches[1])
			y, _ := strconv.Atoi(matches[2])
			vx, _ := strconv.Atoi(matches[3])
			vy, _ := strconv.Atoi(matches[4])
			points = append(points, Point{x, y, vx, vy})
		}
	}
	return points
}

// advance moves all points by t seconds
func advance(points []Point, t int) []Point {
	newPoints := make([]Point, len(points))
	for i, p := range points {
		newPoints[i] = Point{p.x + p.vx*t, p.y + p.vy*t, p.vx, p.vy}
	}
	return newPoints
}

// boundingBox returns min/max x and y
func boundingBox(points []Point) (minX, maxX, minY, maxY int) {
	if len(points) == 0 {
		return 0, 0, 0, 0
	}
	minX, maxX = points[0].x, points[0].x
	minY, maxY = points[0].y, points[0].y
	for _, p := range points[1:] {
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
	return
}

// render prints points as '#' in a grid
func render(points []Point) string {
	minX, maxX, minY, maxY := boundingBox(points)
	width := maxX - minX + 1
	height := maxY - minY + 1

	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	for _, p := range points {
		x := p.x - minX
		y := p.y - minY
		grid[y][x] = '#'
	}

	s := ""
	for _, row := range grid {
		s += string(row) + "\n"
	}
	return s
}
