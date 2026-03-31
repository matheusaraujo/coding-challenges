package main

func part2(puzzleInput []string) any {
	boost := 1

	for {
		army, units, stalemate := simulate(puzzleInput, boost)
		if !stalemate && army == "Immune" {
			return units
		}
		boost++
	}
}
