package main

func part1(puzzleInput []string) any {
	intervals, nums := parseInput(puzzleInput)
	count := 0

	for _, n := range nums {
		for _, i := range intervals {
			if n >= i.Start && n <= i.End {
				count++
				break
			}
		}
	}

	return count
}
