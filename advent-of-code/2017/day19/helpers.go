package main

import (
	"strings"
)

func path(grid []string) (letters string, steps int) {
	x := strings.Index(grid[0], "|")
	y := 0
	dx, dy := 0, 1

	for {
		x += dx
		y += dy
		steps++

		if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[y]) {
			break
		}

		c := grid[y][x]

		if c == ' ' {
			break
		}

		if c >= 'A' && c <= 'Z' {
			letters += string(c)
		}

		if c == '+' {
			if dx == 0 {
				if x > 0 && grid[y][x-1] != ' ' {
					dx, dy = -1, 0
				} else {
					dx, dy = 1, 0
				}
			} else {
				if y > 0 && grid[y-1][x] != ' ' {
					dx, dy = 0, -1
				} else {
					dx, dy = 0, 1
				}
			}
		}
	}

	return
}
