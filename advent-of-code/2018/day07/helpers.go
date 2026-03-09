package main

import (
	"strings"
)

type Edge struct {
	from byte
	to   byte
}

func parseSteps(input []string) ([]byte, map[byte][]byte, map[byte]int) {
	graph := map[byte][]byte{}
	inDegree := map[byte]int{}
	nodes := map[byte]bool{}

	for _, line := range input {
		parts := strings.Split(line, " ")

		from := parts[1][0]
		to := parts[7][0]

		graph[from] = append(graph[from], to)
		inDegree[to]++

		nodes[from] = true
		nodes[to] = true
	}

	allNodes := make([]byte, 0, len(nodes))
	for n := range nodes {
		allNodes = append(allNodes, n)
	}

	return allNodes, graph, inDegree
}
