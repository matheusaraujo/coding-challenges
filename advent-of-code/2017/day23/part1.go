package main

import (
	"strconv"
	"strings"
)

func part1(program []string) any {
	regs := make(map[string]int)
	val := func(arg string) int {
		if arg[0] >= 'a' && arg[0] <= 'h' {
			return regs[arg]
		}
		n, _ := strconv.Atoi(arg)
		return n
	}

	ip := 0
	mulCount := 0
	for ip >= 0 && ip < len(program) {
		parts := strings.Fields(program[ip])
		op := parts[0]
		x := parts[1]
		y := ""
		if len(parts) > 2 {
			y = parts[2]
		}

		switch op {
		case "set":
			regs[x] = val(y)
		case "sub":
			regs[x] -= val(y)
		case "mul":
			regs[x] *= val(y)
			mulCount++
		case "jnz":
			if val(x) != 0 {
				ip += val(y)
				continue
			}
		}
		ip++
	}

	return mulCount
}
