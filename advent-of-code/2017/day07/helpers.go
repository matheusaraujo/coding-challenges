package main

import (
	"strconv"
	"strings"
)

type computer struct {
	weight  int
	holding []string
}

func parseInput(puzzleInput []string) map[string]computer {
	computers := make(map[string]computer)

	for _, line := range puzzleInput {
		parts := strings.Split(line, " -> ")
		left := parts[0]

		fields := strings.Fields(left)
		name := fields[0]
		weightStr := strings.Trim(fields[1], "()")
		weight, _ := strconv.Atoi(weightStr)

		var holding []string
		if len(parts) == 2 {
			holding = strings.Split(parts[1], ", ")
		}

		computers[name] = computer{
			weight:  weight,
			holding: holding,
		}
	}

	return computers
}

func findBottom(computers map[string]computer) string {
	seen := make(map[string]bool)
	for _, comp := range computers {
		for _, child := range comp.holding {
			seen[child] = true
		}
	}

	for name := range computers {
		if !seen[name] {
			return name
		}
	}

	return ""
}
