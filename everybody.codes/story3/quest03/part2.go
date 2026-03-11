package main

func part2(puzzleInput []string) any {

	rules := Rules{
		AllowWeak:    true,
		AllowReplace: false,
	}

	return solveTree(puzzleInput, rules)
}
