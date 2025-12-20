package main

import (
	"strconv"
)

func part2(puzzleInput []string) string {
	operations := make([]rune, 0)
	idx := make([]int, 0)
	n := len(puzzleInput) - 1

	for i, o := range puzzleInput[n] {
		if o != ' ' {
			operations = append(operations, o)
			idx = append(idx, i)
		}
	}
	idx = append(idx, len(puzzleInput[n]))

	result := 0

	for i := 1; i < len(idx); i++ {
		j := idx[i-1]
		k := idx[i]

		column := 0
		if operations[i-1] == '*' {
			column = 1
		}
		for a := j; a < k; a++ {
			m := 1
			x := 0
			for b := len(puzzleInput) - 2; b > -1; b-- {
				if string(puzzleInput[b][a]) != " " {
					x += m * int(puzzleInput[b][a]-'0')
					m *= 10
				}
			}
			if x != 0 {
				if operations[i-1] == '*' {
					column *= x
				} else {
					column += x
				}
			}
		}

		result += column
	}

	return strconv.Itoa(result)
}
