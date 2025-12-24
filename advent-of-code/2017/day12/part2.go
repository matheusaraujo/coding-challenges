package main

func part2(puzzleInput []string) any {
	graph := parseInput(puzzleInput)

	visited := make(map[int]bool)
	groups := 0

	for node := range graph {
		if !visited[node] {
			bfs(graph, node, visited)
			groups++
		}
	}

	return groups
}
