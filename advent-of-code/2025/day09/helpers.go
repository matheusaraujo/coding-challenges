package main

import (
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Line struct {
	p, q Point
}

func parseInput(puzzleInput []string) []Point {
	tiles := make([]Point, len(puzzleInput))
	for i, Line := range puzzleInput {
		parts := strings.Split(Line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		tiles[i] = Point{x, y}
	}
	return tiles
}

func area(a, b Point) int {
	return (abs(a.x-b.x) + 1) * (abs(a.y-b.y) + 1)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
