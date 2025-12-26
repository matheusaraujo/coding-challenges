package main

import (
	"regexp"
	"strconv"
)

// turing machine
func part1(puzzleInput []string) any {
	tape := make(map[int]int)
	cursor := 0
	state, steps := parseInput(puzzleInput)

	for i := 0; i < steps; i++ {
		currentVal := tape[cursor]

		switch state {
		case 'A':
			if currentVal == 0 {
				tape[cursor] = 1
				cursor++
				state = 'B'
			} else {
				tape[cursor] = 0
				cursor--
				state = 'C'
			}
		case 'B':
			if currentVal == 0 {
				tape[cursor] = 1
				cursor--
				state = 'A'
			} else {
				tape[cursor] = 1
				cursor++
				state = 'C'
			}
		case 'C':
			if currentVal == 0 {
				tape[cursor] = 1
				cursor++
				state = 'A'
			} else {
				tape[cursor] = 0
				cursor--
				state = 'D'
			}
		case 'D':
			if currentVal == 0 {
				tape[cursor] = 1
				cursor--
				state = 'E'
			} else {
				tape[cursor] = 1
				cursor--
				state = 'C'
			}
		case 'E':
			if currentVal == 0 {
				tape[cursor] = 1
				cursor++
				state = 'F'
			} else {
				tape[cursor] = 1
				cursor++
				state = 'A'
			}
		case 'F':
			if currentVal == 0 {
				tape[cursor] = 1
				cursor++
				state = 'A'
			} else {
				tape[cursor] = 1
				cursor++
				state = 'E'
			}
		}
	}

	checksum := 0
	for _, val := range tape {
		if val == 1 {
			checksum++
		}
	}

	return checksum

}

func parseInput(puzzleInput []string) (rune, int) {
	stateRe := regexp.MustCompile(`\b[A-Z]\b`)
	state := stateRe.FindString(puzzleInput[0])

	stepsRe := regexp.MustCompile(`\d+`)
	stepsStr := stepsRe.FindString(puzzleInput[1])
	steps, _ := strconv.Atoi(stepsStr)

	return rune(state[0]), steps
}
