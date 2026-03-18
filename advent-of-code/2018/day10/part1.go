package main

import (
	"strings"
)

func part1(puzzleInput []string) any {
	points := parseInput(puzzleInput)

	bestTime := 0
	smallestHeight := 1 << 30

	// simulate seconds until points start spreading again
	for t := 0; ; t++ {
		ps := advance(points, t)
		_, _, minY, maxY := boundingBox(ps)
		height := maxY - minY
		if height < smallestHeight {
			smallestHeight = height
			bestTime = t
		} else {
			// height started increasing, so previous t was best
			break
		}
	}

	// render message at bestTime
	messagePoints := advance(points, bestTime)
	return decodeMessage(render(messagePoints))
}

func decodeMessage(grid string) string {
	lines := strings.Split(strings.TrimSpace(grid), "\n")
	if len(lines) == 0 {
		return ""
	}

	width := len(lines[0])
	var letters []string
	startCol := 0

	for x := 0; x < width; x++ {
		isEmpty := true
		for y := 0; y < len(lines); y++ {
			if lines[y][x] == '#' {
				isEmpty = false
				break
			}
		}

		// If we hit a gap or the end, extract the letter block
		if isEmpty || x == width-1 {
			if x > startCol {
				end := x
				if x == width-1 && !isEmpty {
					end++
				}

				// Extract the sub-grid for one letter
				var charBlock []string
				for _, line := range lines {
					charBlock = append(charBlock, line[startCol:end])
				}
				letters = append(letters, recognize(charBlock))
			}
			startCol = x + 1
		}
	}
	return strings.Join(letters, "")
}

func recognize(block []string) string {
	signature := strings.Join(block, " ")

	patterns := map[string]string{
		"###### #..... #..... #..... #####. #..... #..... #..... #..... #.....": "F",
		"#....# ##...# ##...# #.#..# #.#..# #..#.# #..#.# #...## #...## #....#": "N",
		"#####. #....# #....# #....# #####. #..#.. #...#. #...#. #....# #....#": "R",
		".####. #....# #..... #..... #..... #..### #....# #....# #...## .###.#": "G",
		"#####. #....# #....# #....# #####. #..... #..... #..... #..... #.....": "P",
		"#####. #....# #....# #....# #####. #....# #....# #....# #....# #####.": "B",
		"#....# #....# #....# #....# ###### #....# #....# #....# #....# #....#": "H",
	}

	if char, ok := patterns[signature]; ok {
		return char
	}

	trimmedSignature := strings.TrimSpace(signature)
	for sig, char := range patterns {
		if strings.TrimSpace(sig) == trimmedSignature {
			return char
		}
	}

	return "?"
}
