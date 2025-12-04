package main

import (
	"strings"
)

func part2(puzzleInput []string) interface{} {
	runics, inscriptions := parseInput(puzzleInput)
	result := 0

	for _, insc := range inscriptions {
		m := make(map[int]bool)

		for _, r := range runics {
			matches(insc, r, m)
			matches(insc, reverse(r), m)
		}

		result += len(m)
	}

	return result
}

func matches(insc, pat string, m map[int]bool) {
	i := strings.Index(insc, pat)
	for i != -1 {
		for j := i; j < i+len(pat); j++ {
			m[j] = true
		}
		next := strings.Index(insc[i+1:], pat)
		if next == -1 {
			i = -1
		} else {
			i = i + 1 + next
		}
	}
}
