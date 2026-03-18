package main

func part1(puzzleInput []string) any {
	players, lastMarble := parseInput(puzzleInput[0])
	return playGame(players, lastMarble)
}
