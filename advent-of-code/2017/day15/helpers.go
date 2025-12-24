package main

import (
	"strconv"
	"strings"
)

func parseInput(puzzleInput []string) (int, int) {
	line1 := strings.Split(puzzleInput[0], " ")
	line2 := strings.Split(puzzleInput[1], " ")
	a, _ := strconv.Atoi(line1[4])
	b, _ := strconv.Atoi(line2[4])
	return a, b
}

func solve(a, b int, pairs int, validA func(int) bool, validB func(int) bool) int {
	const factorA = 16807
	const factorB = 48271
	const mod = 2147483647

	count := 0

	for i := 0; i < pairs; i++ {
		for {
			a = (a * factorA) % mod
			if validA(a) {
				break
			}
		}
		for {
			b = (b * factorB) % mod
			if validB(b) {
				break
			}
		}

		if (a & 0xFFFF) == (b & 0xFFFF) {
			count++
		}
	}

	return count
}
