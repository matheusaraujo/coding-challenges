package main

import (
	"strings"
)

func solve(puzzleInput []string, iterations int) any {
	rules := parseInput(puzzleInput)
	grid := []string{
		".#.",
		"..#",
		"###",
	}
	for i := 0; i < iterations; i++ {
		grid = enhance(grid, rules)
	}
	return count(grid)
}

func parseInput(puzzleInput []string) map[string][]string {
	rules := make(map[string][]string)

	for _, line := range puzzleInput {
		parts := strings.Split(line, " => ")
		in := deserialize(parts[0])
		out := deserialize(parts[1])

		for _, v := range variants(in) {
			rules[serialize(v)] = out
		}
	}
	return rules
}

func enhance(grid []string, rules map[string][]string) []string {
	size := len(grid)

	var blockSize int
	if size%2 == 0 {
		blockSize = 2
	} else {
		blockSize = 3
	}

	blocks := size / blockSize
	newBlockSize := blockSize + 1
	newSize := blocks * newBlockSize

	newGrid := make([]string, newSize)
	for i := range newGrid {
		newGrid[i] = ""
	}

	for by := 0; by < blocks; by++ {
		for bx := 0; bx < blocks; bx++ {
			block := make([]string, blockSize)
			for y := 0; y < blockSize; y++ {
				block[y] = grid[by*blockSize+y][bx*blockSize : bx*blockSize+blockSize]
			}

			out := rules[serialize(block)]

			for y := 0; y < newBlockSize; y++ {
				newGrid[by*newBlockSize+y] += out[y]
			}
		}
	}

	return newGrid
}

func serialize(p []string) string {
	return strings.Join(p, "/")
}

func deserialize(p string) []string {
	return strings.Split(p, "/")
}

func variants(p []string) [][]string {
	var result [][]string
	cur := p

	for i := 0; i < 4; i++ {
		result = append(result, cur)
		result = append(result, flip(cur))
		cur = rotate(cur)
	}

	return result
}

func rotate(p []string) []string {
	n := len(p)
	res := make([]string, n)
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = p[n-1-j][i]
		}
		res[i] = string(row)
	}
	return res
}

func flip(p []string) []string {
	res := make([]string, len(p))
	for i, row := range p {
		r := []byte(row)
		for l, h := 0, len(r)-1; l < h; l, h = l+1, h-1 {
			r[l], r[h] = r[h], r[l]
		}
		res[i] = string(r)
	}
	return res
}

func count(grid []string) int {
	cnt := 0
	for _, row := range grid {
		for _, c := range row {
			if c == '#' {
				cnt++
			}
		}
	}

	return cnt
}
