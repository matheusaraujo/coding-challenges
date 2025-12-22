package main

func part1(puzzleInput []string) any {
	plants, id, _, _, _ := parseInput(puzzleInput)
	return plants[id].Energy(id, plants)
}
