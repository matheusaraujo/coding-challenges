package main

func part1(puzzleInput []string) interface{} {
	graph := parseInput(puzzleInput)
	visited := make(map[string]bool)
	return dfs("you", "out", graph, visited)
}

func dfs(node, target string, graph map[string][]string, visited map[string]bool) int {
	if node == target {
		return 1
	}

	if visited[node] {
		return 0
	}

	visited[node] = true
	total := 0

	for _, next := range graph[node] {
		total += dfs(next, target, graph, visited)
	}

	visited[node] = false
	return total
}
