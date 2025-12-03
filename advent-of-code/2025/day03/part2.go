package main

func part2(puzzleInput []string) int {
	result := 0
	for _, bank := range puzzleInput {
		result += joltage12(bank)
	}
	return result
}

func joltage12(bank string) int {
	k := 12
	stack := make([]byte, 0, k)
	drop := len(bank) - k

	for i := 0; i < len(bank); i++ {
		for len(stack) > 0 && drop > 0 && stack[len(stack)-1] < bank[i] {
			stack = stack[:len(stack)-1]
			drop--
		}
		stack = append(stack, bank[i])
	}

	res := stack[:k]

	val := 0
	for _, c := range res {
		val = val*10 + int(c-'0')
	}
	return val
}
