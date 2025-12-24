package main

import (
	"strconv"
	"strings"
)

func parseInput(puzzleInput []string) map[int][]int {
	graph := make(map[int][]int)

	for _, line := range puzzleInput {
		parts := strings.Split(line, " <-> ")
		from, _ := strconv.Atoi(parts[0])

		neighbors := strings.Split(parts[1], ", ")
		for _, n := range neighbors {
			to, _ := strconv.Atoi(n)
			graph[from] = append(graph[from], to)
		}
	}

	return graph
}

func bfs(graph map[int][]int, start int, visited map[int]bool) {
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, n := range graph[cur] {
			if !visited[n] {
				visited[n] = true
				queue = append(queue, n)
			}
		}
	}
}
