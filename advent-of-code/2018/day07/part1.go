package main

import "sort"

func part1(puzzleInput []string) any {
	nodes, graph, inDegree := parseSteps(puzzleInput)

	done := map[byte]bool{}
	result := []byte{}

	for len(result) < len(nodes) {

		available := []byte{}

		for _, n := range nodes {
			if !done[n] && inDegree[n] == 0 {
				available = append(available, n)
			}
		}

		sort.Slice(available, func(i, j int) bool {
			return available[i] < available[j]
		})

		step := available[0]
		result = append(result, step)
		done[step] = true

		for _, next := range graph[step] {
			inDegree[next]--
		}
	}

	return string(result)
}
