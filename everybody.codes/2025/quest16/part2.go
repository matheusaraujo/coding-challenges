package main

func part2(puzzleInput []string) any {
	blocks := parseInput(puzzleInput)
	spell := findSpell(blocks)
	return prod(spell)
}
