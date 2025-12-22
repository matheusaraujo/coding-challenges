package main

import (
	"strconv"
)

type point struct {
	x, y int
}

func part2(puzzleInput []string) any {
	input, _ := strconv.Atoi(puzzleInput[0])
	return spiral(input)
}

func spiral(n int) int {
	values := map[point]int{}
	dirs := []point{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}

	neighbors := []point{
		{1, 0}, {1, 1}, {0, 1}, {-1, 1},
		{-1, 0}, {-1, -1}, {0, -1}, {1, -1},
	}

	pos := point{0, 0}
	values[pos] = 1

	step := 1
	dir := 0

	for {
		for i := 0; i < 2; i++ {
			for s := 0; s < step; s++ {
				pos.x += dirs[dir].x
				pos.y += dirs[dir].y

				sum := 0
				for _, n := range neighbors {
					sum += values[point{pos.x + n.x, pos.y + n.y}]
				}

				if sum > n {
					return sum
				}

				values[pos] = sum
			}
			dir = (dir + 1) % 4
		}
		step++
	}

}
