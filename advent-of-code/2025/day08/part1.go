package main

import (
	"sort"
	"strconv"
)

func part1(puzzleInput []string) string {
	points := parseInput(puzzleInput)
	n := len(points)
	pairs := buildPairs(points)

	K := 1000
	if len(pairs) < K {
		K = len(pairs)
	}

	d := newDSU(n)
	for idx := 0; idx < K; idx++ {
		pr := pairs[idx]
		d.union(pr.i, pr.j)
	}

	rootSizes := make(map[int]int)
	for i := 0; i < n; i++ {
		r := d.find(i)
		rootSizes[r]++
	}

	sizes := make([]int, 0, len(rootSizes))
	for _, s := range rootSizes {
		sizes = append(sizes, s)
	}
	sort.Slice(sizes, func(i, j int) bool { return sizes[i] > sizes[j] })

	prod := 1
	for i := 0; i < 3; i++ {
		if i < len(sizes) {
			prod *= sizes[i]
		}
	}

	return strconv.Itoa(prod)
}
