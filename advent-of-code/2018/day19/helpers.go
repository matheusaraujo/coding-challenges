package main

import (
	"strconv"
	"strings"
)

// Op represents a single instruction
type Op struct {
	name    string
	a, b, c int
}

// runOp executes a single instruction on the registers
func runOp(op Op, regs []int) {
	switch op.name {
	case "addr":
		regs[op.c] = regs[op.a] + regs[op.b]
	case "addi":
		regs[op.c] = regs[op.a] + op.b
	case "mulr":
		regs[op.c] = regs[op.a] * regs[op.b]
	case "muli":
		regs[op.c] = regs[op.a] * op.b
	case "banr":
		regs[op.c] = regs[op.a] & regs[op.b]
	case "bani":
		regs[op.c] = regs[op.a] & op.b
	case "borr":
		regs[op.c] = regs[op.a] | regs[op.b]
	case "bori":
		regs[op.c] = regs[op.a] | op.b
	case "setr":
		regs[op.c] = regs[op.a]
	case "seti":
		regs[op.c] = op.a
	case "gtir":
		if op.a > regs[op.b] {
			regs[op.c] = 1
		} else {
			regs[op.c] = 0
		}
	case "gtri":
		if regs[op.a] > op.b {
			regs[op.c] = 1
		} else {
			regs[op.c] = 0
		}
	case "gtrr":
		if regs[op.a] > regs[op.b] {
			regs[op.c] = 1
		} else {
			regs[op.c] = 0
		}
	case "eqir":
		if op.a == regs[op.b] {
			regs[op.c] = 1
		} else {
			regs[op.c] = 0
		}
	case "eqri":
		if regs[op.a] == op.b {
			regs[op.c] = 1
		} else {
			regs[op.c] = 0
		}
	case "eqrr":
		if regs[op.a] == regs[op.b] {
			regs[op.c] = 1
		} else {
			regs[op.c] = 0
		}
	}
}

// sumDivisors is the "decompiled" logic of the assembly program
func sumDivisors(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}

func parse(input []string) (int, []Op) {
	ipReg, _ := strconv.Atoi(strings.Fields(input[0])[1])
	ops := []Op{}
	for i := 1; i < len(input); i++ {
		f := strings.Fields(input[i])
		a, _ := strconv.Atoi(f[1])
		b, _ := strconv.Atoi(f[2])
		c, _ := strconv.Atoi(f[3])
		ops = append(ops, Op{f[0], a, b, c})
	}
	return ipReg, ops
}
