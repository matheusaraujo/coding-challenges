package main

import (
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func parsePoints(input []string) []Point {
	points := make([]Point, 0, len(input))

	for _, line := range input {
		parts := strings.Split(line, ", ")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])

		points = append(points, Point{x, y})
	}

	return points
}

func manhattan(a, b Point) int {
	dx := a.x - b.x
	if dx < 0 {
		dx = -dx
	}

	dy := a.y - b.y
	if dy < 0 {
		dy = -dy
	}

	return dx + dy
}