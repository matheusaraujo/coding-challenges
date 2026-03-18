package main

import (
	"strconv"
	"strings"
)

// playGame runs the marble game and returns the highest score
func playGame(players int, lastMarble int) int {
	scores := make([]int, players)

	// "Linked list" using arrays
	next := make([]int, lastMarble+1)
	prev := make([]int, lastMarble+1)

	// initial marble 0
	next[0] = 0
	prev[0] = 0

	current := 0

	for marble := 1; marble <= lastMarble; marble++ {
		player := (marble - 1) % players

		if marble%23 == 0 {
			// move 7 CCW
			for i := 0; i < 7; i++ {
				current = prev[current]
			}

			removed := current
			scores[player] += marble + removed

			// remove node
			p := prev[removed]
			n := next[removed]
			next[p] = n
			prev[n] = p

			current = n
		} else {
			// insert between 1 and 2 CW
			a := next[current]
			b := next[a]

			next[a] = marble
			prev[marble] = a

			next[marble] = b
			prev[b] = marble

			current = marble
		}
	}

	// max score
	max := 0
	for _, s := range scores {
		if s > max {
			max = s
		}
	}

	return max
}

// parse input like:
// "10 players; last marble is worth 1618 points"
func parseInput(input string) (int, int) {
	parts := strings.Split(input, " ")
	players, _ := strconv.Atoi(parts[0])
	lastMarble, _ := strconv.Atoi(parts[6])
	return players, lastMarble
}
