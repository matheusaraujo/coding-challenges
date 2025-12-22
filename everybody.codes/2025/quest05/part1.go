package main

func part1(puzzleInput []string) any {
	sword := buildSword(puzzleInput[0])
	return sword.quality
}
