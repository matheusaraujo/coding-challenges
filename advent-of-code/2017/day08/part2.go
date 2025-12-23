package main

func part2(puzzleInput []string) any {
	m, regs := 0, registers{}
	for _, line := range puzzleInput {
		i := parseInstruction(line)
		if i.valid(regs) {
			regs.apply(i)
			m = max(m, regs[i.reg1])
		}
	}
	return m
}
