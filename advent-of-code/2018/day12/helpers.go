package main

import (
	"strings"
)

// parseInputDay12 parses the initial state and rules
func parseInputDay12(lines []string) (map[int]bool, map[string]bool) {
	pots := make(map[int]bool)
	rules := make(map[string]bool)

	// initial state: "initial state: #..#.#..##......###...###"
	for _, line := range lines {
		if strings.HasPrefix(line, "initial state: ") {
			state := strings.TrimPrefix(line, "initial state: ")
			for i, c := range state {
				if c == '#' {
					pots[i] = true
				}
			}
		} else if strings.Contains(line, "=>") {
			parts := strings.Split(line, " => ")
			if len(parts) == 2 && parts[1] == "#" {
				rules[parts[0]] = true
			}
		}
	}

	return pots, rules
}

// nextGeneration computes the next generation of pots
func nextGeneration(current map[int]bool, rules map[string]bool) map[int]bool {
	next := make(map[int]bool)

	// determine range
	min := 0
	max := 0
	for i := range current {
		if i < min {
			min = i
		}
		if i > max {
			max = i
		}
	}

	// extend 2 on each side
	for i := min - 2; i <= max+2; i++ {
		// build pattern string for this pot
		pattern := ""
		for j := i - 2; j <= i+2; j++ {
			if current[j] {
				pattern += "#"
			} else {
				pattern += "."
			}
		}
		if rules[pattern] {
			next[i] = true
		}
	}

	return next
}

// sumPots returns the sum of all pots containing plants
func sumPots(pots map[int]bool) int {
	sum := 0
	for i := range pots {
		sum += i
	}
	return sum
}
