package main

import (
	"strconv"
	"strings"
)

type Point4D struct {
	x, y, z, w int
}

// Manhattan distance: |x1-x2| + |y1-y2| + |z1-z2| + |w1-w2|
func distance(p1, p2 Point4D) int {
	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}
	return abs(p1.x-p2.x) + abs(p1.y-p2.y) + abs(p1.z-p2.z) + abs(p1.w-p2.w)
}

func part1(puzzleInput []string) any {
	var points []Point4D

	// 1. Parse Input
	for _, line := range puzzleInput {
		if line == "" {
			continue
		}
		coords := strings.Split(line, ",")
		p := Point4D{
			x: parseInt(coords[0]),
			y: parseInt(coords[1]),
			z: parseInt(coords[2]),
			w: parseInt(coords[3]),
		}
		points = append(points, p)
	}

	visited := make(map[int]bool)
	constellations := 0

	// 2. Find Connected Components
	for i := 0; i < len(points); i++ {
		if visited[i] {
			continue
		}

		// Found a new constellation
		constellations++
		queue := []int{i}
		visited[i] = true

		// BFS to find all points in this constellation
		for len(queue) > 0 {
			currIdx := queue[0]
			queue = queue[1:]

			for nextIdx := 0; nextIdx < len(points); nextIdx++ {
				if visited[nextIdx] {
					continue
				}

				if distance(points[currIdx], points[nextIdx]) <= 3 {
					visited[nextIdx] = true
					queue = append(queue, nextIdx)
				}
			}
		}
	}

	return constellations
}

func parseInt(s string) int {
	s = strings.TrimSpace(s)
	n, _ := strconv.Atoi(s)
	return n
}
