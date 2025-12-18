package main

import "strings"

func solve(puzzleInput []string, formatWord func(string) string) int {
	count := 0
	for _, line := range puzzleInput {
		if valid(line, formatWord) {
			count++
		}
	}
	return count
}

func valid(passphrase string, formatWord func(string) string) bool {
	seen := make(map[string]bool)
	words := strings.Split(passphrase, " ")
	for _, word := range words {
		w := formatWord(word)
		if seen[w] {
			return false
		}
		seen[w] = true
	}
	return true
}
