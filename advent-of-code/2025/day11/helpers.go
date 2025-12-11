package main

import (
	"strings"
)

func parseInput(input []string) map[string][]string {
	graph := make(map[string][]string)

	for _, line := range input {
		parts := strings.Split(line, ":")
		node := strings.TrimSpace(parts[0])
		neighbors := strings.Fields(parts[1])
		graph[node] = neighbors
	}

	return graph
}
