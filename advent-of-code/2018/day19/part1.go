package main

func part1(puzzleInput []string) any {
	ipReg, ops := parse(puzzleInput)
	regs := make([]int, 6)
	ip := 0

	for ip >= 0 && ip < len(ops) {
		regs[ipReg] = ip
		runOp(ops[ip], regs)
		ip = regs[ipReg] + 1
	}
	return regs[0]
}
