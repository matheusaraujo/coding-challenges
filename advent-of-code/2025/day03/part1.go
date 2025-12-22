package main

func part1(puzzleInput []string) any {
	result := 0
	for _, bank := range puzzleInput {
		m := joltage2(bank)
		result += m
	}
	return result
}

func joltage2(bank string) int {
	md, mu := btoi(bank, 0), btoi(bank, 1)

	for i := 2; i < len(bank); i++ {
		m := btoi(bank, i)
		if mu*10+m > md*10+mu {
			md, mu = mu, m
		} else if md*10+m > md*10+mu {
			mu = m
		}
	}

	return md*10 + mu
}
