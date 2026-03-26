package main

func part2(puzzleInput []string) any {
	i := 3
	for {
		ok, score := solve(puzzleInput, i, true)
		if ok {
			return score
		}
		i++
	}
}
