package main

func part1(puzzleInput []string) any {

	rules := Rules{
		AllowWeak:    false,
		AllowReplace: false,
	}

	return solveTree(puzzleInput, rules)
}
