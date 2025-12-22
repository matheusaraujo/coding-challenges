package main

import (
	"fmt"
)

func part2(puzzleInput []string) any {
	graph := parseInput(puzzleInput)
	required := map[string]bool{
		"fft": false,
		"dac": false,
	}
	memo := make(map[string]map[string]int)
	return dfsMemo("svr", "out", graph, required, memo)
}

func dfsMemo(node, target string, graph map[string][]string, required map[string]bool, memo map[string]map[string]int) int {
	if _, ok := required[node]; ok {
		required[node] = true
	}

	allSeen := visitedAll(required)

	if node == target {
		if allSeen {
			return 1
		}
		return 0
	}

	seenKey := fmt.Sprintf("%v", required)

	if _, ok := memo[node]; ok {
		if val, ok2 := memo[node][seenKey]; ok2 {
			return val
		}
	} else {
		memo[node] = make(map[string]int)
	}

	total := 0
	for _, next := range graph[node] {
		total += dfsMemo(next, target, graph, clone(required), memo)
	}

	memo[node][seenKey] = total
	return total
}

func visitedAll(required map[string]bool) bool {
	for _, v := range required {
		if !v {
			return false
		}
	}
	return true
}

func clone(src map[string]bool) map[string]bool {
	dst := make(map[string]bool, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}
