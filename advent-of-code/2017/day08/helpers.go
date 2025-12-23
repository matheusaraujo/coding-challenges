package main

import (
	"strconv"
	"strings"
)

type instruction struct {
	reg1   string
	op     string
	opArg  int
	reg2   string
	cmp    string
	cmpArg int
}

type registers map[string]int

func (i instruction) valid(regs registers) bool {
	val := regs[i.reg2]
	switch i.cmp {
	case ">":
		return val > i.cmpArg
	case "<":
		return val < i.cmpArg
	case ">=":
		return val >= i.cmpArg
	case "<=":
		return val <= i.cmpArg
	case "==":
		return val == i.cmpArg
	case "!=":
		return val != i.cmpArg
	default:
		return false
	}
}

func (regs registers) apply(i instruction) {
	switch i.op {
	case "inc":
		regs[i.reg1] += i.opArg
	case "dec":
		regs[i.reg1] -= i.opArg
	}
}

func (regs registers) max() int {
	m := 0
	for _, v := range regs {
		m = max(m, v)
	}
	return m
}

func parseInstruction(line string) instruction {
	parts := strings.Split(line, " ")
	opArg, _ := strconv.Atoi(parts[2])
	cmpArg, _ := strconv.Atoi(parts[6])

	return instruction{
		reg1:   parts[0],
		op:     parts[1],
		opArg:  opArg,
		reg2:   parts[4],
		cmp:    parts[5],
		cmpArg: cmpArg,
	}
}
