package main

import (
	"strconv"
	"strings"
)

const SHAPES = 6

func part1(puzzleInput []string) any {
	shapes, regions := parseInput(puzzleInput)
	return solve(shapes, regions)
}

func parseInput(puzzleInput []string) ([]int, [][]int) {
	shapes := make([]int, SHAPES)
	regions := make([][]int, 0)

	for i := 0; i < SHAPES; i++ {
		lines := puzzleInput[i*5+1 : i*5+4]
		shapes[i] = strings.Count(lines[0], "#") + strings.Count(lines[1], "#") + strings.Count(lines[2], "#")
	}

	for i := SHAPES * 5; i < len(puzzleInput); i++ {
		parts := strings.Split(puzzleInput[i], ":")
		dims := strings.Split(parts[0], "x")
		reqs := strings.Split(parts[1], " ")[1:]

		region := make([]int, 1+len(reqs))
		region[0] = atoi(dims[0]) * atoi(dims[1])

		for j := 0; j < len(reqs); j++ {
			region[j+1] = atoi(reqs[j])
		}

		regions = append(regions, region)
	}

	return shapes, regions
}

func solve(shapes []int, regions [][]int) int {
	result := 0

	for i := 0; i < len(regions); i++ {
		s := 0
		for j := 0; j < SHAPES; j++ {
			s += shapes[j] * regions[i][j+1]
		}
		if s < regions[i][0] {
			result++
		}
	}

	return result
}

func atoi(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}
