package main

import (
	"container/heap"
)

func part2(puzzleInput []string) any {
	depth, target := parseInput(puzzleInput)

	// Extend grid beyond target (important!)
	maxX := target.x + 50
	maxY := target.y + 50

	grid := buildGrid(depth, target, maxX, maxY)

	start := State{0, 0, Torch, 0}

	visited := make(map[[3]int]int)

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, start)

	dirs := []struct{ dx, dy int }{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1},
	}

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(State)

		key := [3]int{cur.x, cur.y, cur.tool}

		if t, ok := visited[key]; ok && t <= cur.time {
			continue
		}
		visited[key] = cur.time

		// Goal condition
		if cur.x == target.x && cur.y == target.y && cur.tool == Torch {
			return cur.time
		}

		// Try switching tools
		for _, t := range validTools(grid[cur.y][cur.x]) {
			if t != cur.tool {
				heap.Push(pq, State{
					cur.x, cur.y, t, cur.time + 7,
				})
			}
		}

		// Move
		for _, d := range dirs {
			nx, ny := cur.x+d.dx, cur.y+d.dy

			if nx < 0 || ny < 0 || nx > maxX || ny > maxY {
				continue
			}

			if allows(grid[ny][nx], cur.tool) {
				heap.Push(pq, State{
					nx, ny, cur.tool, cur.time + 1,
				})
			}
		}
	}

	return -1
}
