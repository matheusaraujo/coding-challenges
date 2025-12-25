package main

import "sort"

func part1(puzzleInput []string) any {
	particles := parseInput(puzzleInput)

	sort.Slice(particles, func(i, j int) bool {
		if manhattan(particles[i].a) != manhattan(particles[j].a) {
			return manhattan(particles[i].a) < manhattan(particles[j].a)
		}
		if manhattan(particles[i].v) != manhattan(particles[j].v) {
			return manhattan(particles[i].v) < manhattan(particles[j].v)
		}
		return manhattan(particles[i].p) < manhattan(particles[j].p)
	})

	return particles[0].id
}

func manhattan(v [3]int) int {
	sum := 0
	for _, x := range v {
		if x < 0 {
			sum -= x
		} else {
			sum += x
		}
	}
	return sum
}
