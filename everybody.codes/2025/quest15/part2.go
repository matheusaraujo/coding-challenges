package main

import (
	"strconv"
)

func part2(puzzleInput []string) string {
	m, e, minX, maxX, minY, maxY := buildMap(puzzleInput)
	return strconv.Itoa(bfs(m, ORIGIN, e, minX, maxX, minY, maxY))
}
