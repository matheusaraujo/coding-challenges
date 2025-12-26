package main

import (
	"strconv"
	"strings"
)

type component struct {
	a, b int
}

func parseInput(input []string) []component {
	comps := make([]component, len(input))
	for i, line := range input {
		parts := strings.Split(line, "/")
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		comps[i] = component{a, b}
	}
	return comps
}
