package main

func part1(puzzleInput []string) any {
	prog, instruction := parseProgram(puzzleInput)

	resetRegisters()

	prev := initPrevious(15000)
	i := 2
	answer := prog[28].args[0]

	for {
		ip := int(*instruction)

		if ip == 28 {
			if i == 2 {
				return registers[answer]
			}

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

	return nil
}
