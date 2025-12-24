package main

type Program struct {
	id        int
	ptr       int
	regs      map[string]int
	queue     []int
	waiting   bool
	sendCount int
}

func part2(puzzleInput []string) any {
	p0 := Program{
		id:   0,
		regs: map[string]int{"p": 0},
	}
	p1 := Program{
		id:   1,
		regs: map[string]int{"p": 1},
	}

	for {
		step(&p0, puzzleInput, &p1.queue)
		step(&p1, puzzleInput, &p0.queue)

		if p0.waiting && p1.waiting {
			return p1.sendCount
		}
	}
}

func step(p *Program, insts []string, out *[]int) {
	if p.ptr < 0 || p.ptr >= len(insts) {
		p.waiting = true
		return
	}

	op, x, y, offset := inst(insts[p.ptr])

	switch op {
	case "snd":
		v := val(x, p.regs)
		*out = append(*out, v)
		p.sendCount++
	case "set":
		p.regs[x] = val(y, p.regs)
	case "add":
		p.regs[x] += val(y, p.regs)
	case "mul":
		p.regs[x] *= val(y, p.regs)
	case "mod":
		p.regs[x] %= val(y, p.regs)
	case "rcv":
		if len(p.queue) == 0 {
			p.waiting = true
			offset = 0
		} else {
			p.regs[x] = p.queue[0]
			p.queue = p.queue[1:]
			p.waiting = false
		}
	case "jgz":
		if val(x, p.regs) > 0 {
			offset = val(y, p.regs)
		}
	}
	p.ptr += offset
}
