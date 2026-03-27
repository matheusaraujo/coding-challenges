package main

func part2(puzzleInput []string) any {
	ipReg, ops := parse(puzzleInput)
	regs := make([]int, 6)
	regs[0] = 1 // Part 2 starts with reg 0 as 1
	ip := 0

	// We run the simulation until the "Target Number" is generated in a register.
	// In almost all inputs, this happens within the first 100-200 cycles.
	// Usually, the target number is the largest value in the registers once
	// the instruction pointer starts looping between small values.
	for i := 0; i < 500; i++ {
		regs[ipReg] = ip
		runOp(ops[ip], regs)
		ip = regs[ipReg] + 1
	}

	// Find the large number N that the program is trying to factorize
	maxN := 0
	for _, v := range regs {
		if v > maxN {
			maxN = v
		}
	}

	return sumDivisors(maxN)
}
