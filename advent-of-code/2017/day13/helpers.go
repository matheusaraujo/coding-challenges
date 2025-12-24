package main

import (
	"strconv"
	"strings"
)

func parse(line string) (int, int) {
	parts := strings.Split(line, ": ")
	depth, _ := strconv.Atoi(parts[0])
	rng, _ := strconv.Atoi(parts[1])
	return depth, rng
}
