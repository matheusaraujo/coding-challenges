package main

import (
	"sort"
	"strconv"
	"strings"
)

func parseInput(puzzleInput []string) [][]int {
	var input [][]int
	for _, line := range puzzleInput {
		ns := strings.Split(line, "\t")
		n := make([]int, len(ns))
		for i, s := range ns {
			n[i], _ = strconv.Atoi(s)
		}
		sort.Ints(n)
		input = append(input, n)
	}
	return input
}
