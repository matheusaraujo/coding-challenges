// based on https://www.reddit.com/r/adventofcode/comments/a86jgt/comment/ec8i366/

package main

import (
	"fmt"
)

type OpFunc func(a, b, c int)

type Op struct {
	f    OpFunc
	args [3]int
}

var registers [6]int64

// ---------------- Parsing ----------------

func parseProgram(input []string) ([]Op, *int64) {
	var inline string
	var ipReg int

	fmt.Sscanf(input[0], "%s %d", &inline, &ipReg)
	instruction := &registers[ipReg]

	var prog []Op

	for _, line := range input[1:] {
		var name string
		var a, b, c int
		fmt.Sscanf(line, "%s %d %d %d", &name, &a, &b, &c)

		op := Op{args: [3]int{a, b, c}}

		switch name {
		case "addr":
			op.f = addr
		case "addi":
			op.f = addi
		case "mulr":
			op.f = mulr
		case "muli":
			op.f = muli
		case "banr":
			op.f = banr
		case "bani":
			op.f = bani
		case "borr":
			op.f = borr
		case "bori":
			op.f = bori
		case "setr":
			op.f = setr
		case "seti":
			op.f = seti
		case "gtir":
			op.f = gtir
		case "gtri":
			op.f = gtri
		case "gtrr":
			op.f = gtrr
		case "eqir":
			op.f = eqir
		case "eqri":
			op.f = eqri
		case "eqrr":
			op.f = eqrr
		}

		prog = append(prog, op)
	}

	return prog, instruction
}

// ---------------- Execution helpers ----------------

func execOp(op Op) {
	op.f(op.args[0], op.args[1], op.args[2])
}

func resetRegisters() {
	for i := range registers {
		registers[i] = 0
	}
}

// ---------------- Previous state tracking ----------------

func initPrevious(size int) [][]int64 {
	prev := make([][]int64, 6)
	for i := range prev {
		prev[i] = make([]int64, size)
	}
	return prev
}

func extend(prev [][]int64) [][]int64 {
	newSize := len(prev[0]) + 1000
	newPrev := make([][]int64, 6)
	for i := range newPrev {
		newPrev[i] = make([]int64, newSize)
		copy(newPrev[i], prev[i])
	}
	return newPrev
}

func seen(prev [][]int64, regs [6]int64, upto int) bool {
	for j := 1; j < upto; j++ {
		match := true
		for k := 0; k < 6; k++ {
			if prev[k][j] != regs[k] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

func store(prev [][]int64, regs [6]int64, idx int) {
	for k := 0; k < 6; k++ {
		prev[k][idx] = regs[k]
	}
}

func duplicateInColumn(prev [][]int64, col int, idx int) bool {
	for j := 1; j < idx; j++ {
		if prev[col][j] == prev[col][idx] {
			return true
		}
	}
	return false
}

// ---------------- Operations ----------------

func addr(a, b, c int) { registers[c] = registers[a] + registers[b] }
func addi(a, b, c int) { registers[c] = registers[a] + int64(b) }
func mulr(a, b, c int) { registers[c] = registers[a] * registers[b] }
func muli(a, b, c int) { registers[c] = registers[a] * int64(b) }
func banr(a, b, c int) { registers[c] = registers[a] & registers[b] }
func bani(a, b, c int) { registers[c] = registers[a] & int64(b) }
func borr(a, b, c int) { registers[c] = registers[a] | registers[b] }
func bori(a, b, c int) { registers[c] = registers[a] | int64(b) }
func setr(a, b, c int) { registers[c] = registers[a] }
func seti(a, b, c int) { registers[c] = int64(a) }

func gtir(a, b, c int) {
	if int64(a) > registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func gtri(a, b, c int) {
	if registers[a] > int64(b) {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func gtrr(a, b, c int) {
	if registers[a] > registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func eqir(a, b, c int) {
	if int64(a) == registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func eqri(a, b, c int) {
	if registers[a] == int64(b) {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func eqrr(a, b, c int) {
	if registers[a] == registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}
