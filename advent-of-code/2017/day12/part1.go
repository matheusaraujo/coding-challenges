package main

func part1(puzzleInput []string) any {
	graph := parseInput(puzzleInput)
	visited := make(map[int]bool)
	bfs(graph, 0, visited)
	return len(visited)
}
