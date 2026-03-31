package main

import (
	"fmt"
)

type Bot struct {
	x, y, z int
	r       int
}

func parseBots(input []string) []Bot {
	bots := make([]Bot, 0, len(input))

	for _, line := range input {
		var x, y, z, r int
		fmt.Sscanf(line, "pos=<%d,%d,%d>, r=%d", &x, &y, &z, &r)
		bots = append(bots, Bot{x, y, z, r})
	}

	return bots
}

func manhattan(x1, y1, z1, x2, y2, z2 int) int {
	return abs(x1-x2) + abs(y1-y2) + abs(z1-z2)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
