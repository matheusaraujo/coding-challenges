package main

import (
	"strconv"
)

func part1(puzzleInput []string) string {
	intervals, nums := parseInput(puzzleInput)
	count := 0

	for _, n := range nums {
		for _, i := range intervals {
			if n >= i.Start && n <= i.End {
				count++
				break
			}
		}
	}

	return strconv.Itoa(count)
}
