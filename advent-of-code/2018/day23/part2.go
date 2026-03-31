package main

import (
	"container/heap"
)

type Cube struct {
	x, y, z int // min corner
	size    int
	bots    int
	dist    int
}

// Priority queue
type PQ []Cube

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	// Max bots first
	if pq[i].bots != pq[j].bots {
		return pq[i].bots > pq[j].bots
	}
	// Then closest to origin
	if pq[i].dist != pq[j].dist {
		return pq[i].dist < pq[j].dist
	}
	// Then smaller cubes
	return pq[i].size < pq[j].size
}

func (pq PQ) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PQ) Push(x any) {
	*pq = append(*pq, x.(Cube))
}

func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

// Distance from cube to origin
func cubeDist(x, y, z, size int) int {
	dx := max(0, abs(x)) + max(0, abs(x+size-1))
	dy := max(0, abs(y)) + max(0, abs(y+size-1))
	dz := max(0, abs(z)) + max(0, abs(z+size-1))
	return dx + dy + dz
}

func botsInCube(bots []Bot, c Cube) int {
	count := 0
	for _, b := range bots {
		d := 0

		if b.x < c.x {
			d += c.x - b.x
		} else if b.x > c.x+c.size-1 {
			d += b.x - (c.x + c.size - 1)
		}

		if b.y < c.y {
			d += c.y - b.y
		} else if b.y > c.y+c.size-1 {
			d += b.y - (c.y + c.size - 1)
		}

		if b.z < c.z {
			d += c.z - b.z
		} else if b.z > c.z+c.size-1 {
			d += b.z - (c.z + c.size - 1)
		}

		if d <= b.r {
			count++
		}
	}
	return count
}

func part2(puzzleInput []string) any {
	bots := parseBots(puzzleInput)

	// Bounding box
	minX, maxX := bots[0].x, bots[0].x
	minY, maxY := bots[0].y, bots[0].y
	minZ, maxZ := bots[0].z, bots[0].z

	for _, b := range bots {
		if b.x < minX {
			minX = b.x
		}
		if b.x > maxX {
			maxX = b.x
		}
		if b.y < minY {
			minY = b.y
		}
		if b.y > maxY {
			maxY = b.y
		}
		if b.z < minZ {
			minZ = b.z
		}
		if b.z > maxZ {
			maxZ = b.z
		}
	}

	// Find power-of-two cube size
	size := 1
	for size < max(maxX-minX, max(maxY-minY, maxZ-minZ)) {
		size *= 2
	}

	pq := &PQ{}
	heap.Init(pq)

	start := Cube{minX, minY, minZ, size, 0, 0}
	start.bots = botsInCube(bots, start)
	start.dist = manhattan(0, 0, 0, start.x, start.y, start.z)

	heap.Push(pq, start)

	for pq.Len() > 0 {
		c := heap.Pop(pq).(Cube)

		if c.size == 1 {
			return c.dist
		}

		half := c.size / 2

		for dx := 0; dx < 2; dx++ {
			for dy := 0; dy < 2; dy++ {
				for dz := 0; dz < 2; dz++ {
					nc := Cube{
						x:    c.x + dx*half,
						y:    c.y + dy*half,
						z:    c.z + dz*half,
						size: half,
					}
					nc.bots = botsInCube(bots, nc)
					nc.dist = manhattan(0, 0, 0, nc.x, nc.y, nc.z)

					heap.Push(pq, nc)
				}
			}
		}
	}

	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
