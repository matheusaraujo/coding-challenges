package main

import (
	"strconv"
	"strings"
)

func parseInput(puzzleInput []string) ([]string, []rune) {
	moves := strings.Split(puzzleInput[0], ",")

	dancers := make([]rune, 16)
	for i := 0; i < 16; i++ {
		dancers[i] = rune('a' + i)
	}
	return moves, dancers
}

func dance(dancers []rune, moves []string) []rune {
	for _, move := range moves {
		switch move[0] {
		case 's':
			dancers = spin(move, dancers)
		case 'x':
			dancers = exchange(move, dancers)
		case 'p':
			dancers = partner(move, dancers)
		}
	}
	return dancers
}

func spin(move string, dancers []rune) []rune {
	x, _ := strconv.Atoi(move[1:])
	n := len(dancers)
	dancers = append(dancers[n-x:], dancers[:n-x]...)
	return dancers
}

func exchange(move string, dancers []rune) []rune {
	parts := strings.Split(move[1:], "/")
	a, _ := strconv.Atoi(parts[0])
	b, _ := strconv.Atoi(parts[1])
	dancers[a], dancers[b] = dancers[b], dancers[a]
	return dancers
}

func partner(move string, dancers []rune) []rune {
	a := rune(move[1])
	b := rune(move[3])
	var ia, ib int
	for i, d := range dancers {
		if d == a {
			ia = i
		} else if d == b {
			ib = i
		}
	}
	dancers[ia], dancers[ib] = dancers[ib], dancers[ia]
	return dancers
}
