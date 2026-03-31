package main

func part2(puzzleInput []string) any {
	prog, instruction := parseProgram(puzzleInput)

	resetRegisters()

	prev := initPrevious(15000)
	i := 2
	answer := prog[28].args[0]

	for {
		ip := int(*instruction)

		if ip == 28 {
			if i >= len(prev[0]) {
				prev = extend(prev)
			}

			if seen(prev, registers, i) {
				break
			}

			store(prev, registers, i)
			i++
		}

		execOp(prog[ip])
		*instruction++
	}

	// find last unique value
	for {
		i--
		if duplicateInColumn(prev, answer, i) {
			continue
		}
		return prev[answer][i]
	}
}
