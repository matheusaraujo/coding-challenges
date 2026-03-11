package main

func part3(puzzleInput []string) any {

	rules := Rules{
		AllowWeak:    true,
		AllowReplace: true,
	}

	return solveTree(puzzleInput, rules)
}
