package main

func part2(puzzleInput []string) any {
	input := puzzleInput[0]

	target := make([]int, len(input))
	for i, ch := range input {
		target[i] = int(ch - '0')
	}

	recipes := []int{3, 7}
	e1, e2 := 0, 1

	for {
		prevLen := len(recipes)
		step(&recipes, &e1, &e2)

		// check last positions (important: can grow by 2)
		if endsWith(recipes, target) {
			return len(recipes) - len(target)
		}
		if endsWithOffset(recipes, target, 1) {
			return len(recipes) - len(target) - 1
		}

		// (optional micro-optimization: skip unnecessary checks)
		if len(recipes) == prevLen {
			continue
		}
	}
}
