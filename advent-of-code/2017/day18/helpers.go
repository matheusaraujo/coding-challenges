package main

import (
	"strconv"
	"strings"
)

func inst(inst string) (op, x, y string, offset int) {
	parts := strings.Fields(inst)

	op = parts[0]
	x = parts[1]
	offset = 1

	if len(parts) > 2 {
		y = parts[2]
	} else {
		y = ""
	}

	return
}

func val(x string, regs map[string]int) int {
	if x[0] >= 'a' && x[0] <= 'z' {
		return regs[x]
	}
	v, _ := strconv.Atoi(x)
	return v
}
