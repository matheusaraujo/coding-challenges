package main

func part2(puzzleInput []string) any {
	m, e, minX, maxX, minY, maxY := buildMap(puzzleInput)
	return bfs(m, ORIGIN, e, minX, maxX, minY, maxY)
}
