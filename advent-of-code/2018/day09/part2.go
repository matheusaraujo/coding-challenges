package main

func part2(puzzleInput []string) any {
	players, lastMarble := parseInput(puzzleInput[0])
	return playGame(players, lastMarble*100)
}
