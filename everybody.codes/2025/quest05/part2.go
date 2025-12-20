package main

import (
	"math"
	"strconv"
)

func part2(puzzleInput []string) string {
	mi, ma := math.MaxInt, math.MinInt
	for i := 0; i < len(puzzleInput); i++ {
		sword := buildSword(puzzleInput[i])
		mi = min(mi, sword.quality)
		ma = max(ma, sword.quality)
	}
	return strconv.Itoa(ma - mi)
}
