package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solve(puzzleInput []string) (int, int) {
	banks := parseInput(puzzleInput)
	seen := make(map[string]int)
	cycles := 0

	for {
		state := fmt.Sprint(banks)

		if firstSeenAt, exists := seen[state]; exists {
			loopSize := cycles - firstSeenAt
			return cycles, loopSize
		}

		seen[state] = cycles

		maxVal := -1
		maxIdx := -1
		for i, val := range banks {
			if val > maxVal {
				maxVal = val
				maxIdx = i
			}
		}

		blocks := banks[maxIdx]
		banks[maxIdx] = 0

		curr := maxIdx
		for blocks > 0 {
			curr = (curr + 1) % len(banks)
			banks[curr]++
			blocks--
		}

		cycles++
	}
}

func parseInput(puzzleInput []string) []int {
	fields := strings.Fields(puzzleInput[0])
	banks := make([]int, len(fields))
	for i, f := range fields {
		banks[i], _ = strconv.Atoi(f)
	}
	return banks
}
