package main

func part1(puzzleInput []string) any {
	regs := registers{}
	for _, line := range puzzleInput {
		i := parseInstruction(line)
		if i.valid(regs) {
			regs.apply(i)
		}
	}
	return regs.max()
}
