package main

import (
	"sort"
	"strconv"
	"strings"
)

type point struct {
	x, y, z int
}

type pair struct {
	i, j  int
	dist2 int
}

type dsu struct {
	parent []int
	size   []int
}

func newDSU(n int) *dsu {
	p := make([]int, n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
		s[i] = 1
	}
	return &dsu{parent: p, size: s}
}

func (d *dsu) find(a int) int {
	for d.parent[a] != a {
		d.parent[a] = d.parent[d.parent[a]]
		a = d.parent[a]
	}
	return a
}

func (d *dsu) union(a, b int) bool {
	ra := d.find(a)
	rb := d.find(b)
	if ra == rb {
		return false
	}
	if d.size[ra] < d.size[rb] {
		ra, rb = rb, ra
	}
	d.parent[rb] = ra
	d.size[ra] += d.size[rb]
	return true
}

func squaredDist(a, b point) int {
	dx := a.x - b.x
	dy := a.y - b.y
	dz := a.z - b.z
	return dx*dx + dy*dy + dz*dz
}

func parseInput(puzzleInput []string) []point {
	points := make([]point, 0, len(puzzleInput))

	for _, line := range puzzleInput {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		points = append(points, point{x, y, z})
	}

	return points
}

func buildPairs(points []point) []pair {
	n := len(points)
	pairs := make([]pair, 0, n*(n-1)/2)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			pairs = append(pairs, pair{
				i:     i,
				j:     j,
				dist2: squaredDist(points[i], points[j]),
			})
		}
	}

	sort.Slice(pairs, func(a, b int) bool {
		if pairs[a].dist2 != pairs[b].dist2 {
			return pairs[a].dist2 < pairs[b].dist2
		}
		if pairs[a].i != pairs[b].i {
			return pairs[a].i < pairs[b].i
		}
		return pairs[a].j < pairs[b].j
	})

	return pairs
}
