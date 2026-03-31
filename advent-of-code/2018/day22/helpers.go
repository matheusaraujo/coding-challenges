package main

import (
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Region int

const (
	Rocky Region = iota
	Wet
	Narrow
)

const (
	None = iota
	Torch
	Gear
)

type State struct {
	x, y int
	tool int
	time int
}

// Priority Queue for Dijkstra
type PriorityQueue []State

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].time < pq[j].time
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(State))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

// Parse input
func parseInput(input []string) (int, Point) {
	depthStr := strings.Split(input[0], ": ")[1]
	targetStr := strings.Split(input[1], ": ")[1]

	depth, _ := strconv.Atoi(depthStr)
	parts := strings.Split(targetStr, ",")

	tx, _ := strconv.Atoi(parts[0])
	ty, _ := strconv.Atoi(parts[1])

	return depth, Point{tx, ty}
}

// Compute region types with memoization
func buildGrid(depth int, target Point, maxX, maxY int) [][]Region {
	erosion := make([][]int, maxY+1)
	grid := make([][]Region, maxY+1)

	for y := 0; y <= maxY; y++ {
		erosion[y] = make([]int, maxX+1)
		grid[y] = make([]Region, maxX+1)
		for x := 0; x <= maxX; x++ {
			var geo int

			if (x == 0 && y == 0) || (x == target.x && y == target.y) {
				geo = 0
			} else if y == 0 {
				geo = x * 16807
			} else if x == 0 {
				geo = y * 48271
			} else {
				geo = erosion[y][x-1] * erosion[y-1][x]
			}

			erosion[y][x] = (geo + depth) % 20183
			grid[y][x] = Region(erosion[y][x] % 3)
		}
	}

	return grid
}

// Valid tools per region
func validTools(r Region) []int {
	switch r {
	case Rocky:
		return []int{Torch, Gear}
	case Wet:
		return []int{Gear, None}
	case Narrow:
		return []int{Torch, None}
	}
	return nil
}

// Check if tool allowed
func allows(r Region, tool int) bool {
	for _, t := range validTools(r) {
		if t == tool {
			return true
		}
	}
	return false
}
