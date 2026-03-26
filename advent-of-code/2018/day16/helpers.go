package main

import (
	"regexp"
	"strconv"
	"strings"
)

type Sample struct {
	Before [4]int
	Inst   [4]int
	After  [4]int
}

type OpFunc func(r [4]int, inst [4]int) [4]int

var Opcodes = map[string]OpFunc{
	"addr": func(r [4]int, i [4]int) [4]int { r[i[3]] = r[i[1]] + r[i[2]]; return r },
	"addi": func(r [4]int, i [4]int) [4]int { r[i[3]] = r[i[1]] + i[2]; return r },
	"mulr": func(r [4]int, i [4]int) [4]int { r[i[3]] = r[i[1]] * r[i[2]]; return r },
	"muli": func(r [4]int, i [4]int) [4]int { r[i[3]] = r[i[1]] * i[2]; return r },
	"banr": func(r [4]int, i [4]int) [4]int { r[i[3]] = r[i[1]] & r[i[2]]; return r },
	"bani": func(r [4]int, i [4]int) [4]int { r[i[3]] = r[i[1]] & i[2]; return r },
	"borr": func(r [4]int, i [4]int) [4]int { r[i[3]] = r[i[1]] | r[i[2]]; return r },
	"bori": func(r [4]int, i [4]int) [4]int { r[i[3]] = r[i[1]] | i[2]; return r },
	"setr": func(r [4]int, i [4]int) [4]int { r[i[3]] = r[i[1]]; return r },
	"seti": func(r [4]int, i [4]int) [4]int { r[i[3]] = i[1]; return r },
	"gtir": func(r [4]int, i [4]int) [4]int {
		if i[1] > r[i[2]] {
			r[i[3]] = 1
		} else {
			r[i[3]] = 0
		}
		return r
	},
	"gtri": func(r [4]int, i [4]int) [4]int {
		if r[i[1]] > i[2] {
			r[i[3]] = 1
		} else {
			r[i[3]] = 0
		}
		return r
	},
	"gtrr": func(r [4]int, i [4]int) [4]int {
		if r[i[1]] > r[i[2]] {
			r[i[3]] = 1
		} else {
			r[i[3]] = 0
		}
		return r
	},
	"eqir": func(r [4]int, i [4]int) [4]int {
		if i[1] == r[i[2]] {
			r[i[3]] = 1
		} else {
			r[i[3]] = 0
		}
		return r
	},
	"eqri": func(r [4]int, i [4]int) [4]int {
		if r[i[1]] == i[2] {
			r[i[3]] = 1
		} else {
			r[i[3]] = 0
		}
		return r
	},
	"eqrr": func(r [4]int, i [4]int) [4]int {
		if r[i[1]] == r[i[2]] {
			r[i[3]] = 1
		} else {
			r[i[3]] = 0
		}
		return r
	},
}

func parseInput(lines []string) ([]Sample, [][4]int) {
	var samples []Sample
	var program [][4]int

	reNum := regexp.MustCompile(`\d+`)

	i := 0
	for i < len(lines) {
		line := lines[i]
		if strings.HasPrefix(line, "Before:") {
			var before, inst, after [4]int
			bMatch := reNum.FindAllString(lines[i], -1)
			iMatch := reNum.FindAllString(lines[i+1], -1)
			aMatch := reNum.FindAllString(lines[i+2], -1)

			for j := 0; j < 4; j++ {
				before[j], _ = strconv.Atoi(bMatch[j])
				inst[j], _ = strconv.Atoi(iMatch[j])
				after[j], _ = strconv.Atoi(aMatch[j])
			}
			samples = append(samples, Sample{before, inst, after})
			i += 3
		} else if line == "" {
			i++
		} else {
			var inst [4]int
			match := reNum.FindAllString(line, -1)
			if len(match) == 4 {
				for j := 0; j < 4; j++ {
					inst[j], _ = strconv.Atoi(match[j])
				}
				program = append(program, inst)
			}
			i++
		}
	}
	return samples, program
}
