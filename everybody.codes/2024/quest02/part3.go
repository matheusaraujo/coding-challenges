package main

import (
	"strconv"
)

type pos struct {
	i, j int
}

func part3(puzzleInput []string) string {
	runics, inscriptions := parseInput(puzzleInput)

	m := make(map[pos](bool))

	for i := 0; i < len(inscriptions); i++ {
		for j := 0; j < len(inscriptions[i]); j++ {
			for _, r := range runics {
				for d := 0; d < 4; d++ {
					match2d(inscriptions, i, j, r, d, m)
				}
			}
		}
	}
	return strconv.Itoa(len(m))
}

func match2d(inscriptions []string, i, j int, r string, d int, m map[pos]bool) {
	rows := len(inscriptions)
	cols := len(inscriptions[0])
	n := len(r)

	di := []int{0, 0, 1, -1}
	dj := []int{1, -1, 0, 0}

	ci, cj := i, j

	for k := 0; k < n; k++ {
		ri := ci
		rj := (cj + cols) % cols
		if di[d] != 0 {
			if ci < 0 || ci >= rows {
				return
			}
		}
		if inscriptions[ri][rj] != r[k] {
			return
		}
		ci += di[d]
		cj += dj[d]
	}

	ci, cj = i, j
	for k := 0; k < n; k++ {
		ri := ci
		rj := (cj + cols) % cols
		m[pos{ri, rj}] = true

		ci += di[d]
		cj += dj[d]
	}

	return
}
