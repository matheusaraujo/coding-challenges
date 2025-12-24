package main

func part1(puzzleInput []string) any {
	ptr, sound := 0, 0
	regs := map[string]int{}

	for ptr >= 0 && ptr < len(puzzleInput) {
		op, x, y, offset := inst(puzzleInput[ptr])
		switch op {
		case "snd":
			sound = val(x, regs)
		case "set":
			regs[x] = val(y, regs)
		case "add":
			regs[x] = val(x, regs) + val(y, regs)
		case "mul":
			regs[x] = val(x, regs) * val(y, regs)
		case "mod":
			regs[x] = val(x, regs) % val(y, regs)
		case "rcv":
			if val(x, regs) != 0 {
				return sound
			}
		case "jgz":
			if val(x, regs) > 0 {
				offset = val(y, regs)
			}
		}
		ptr += offset
	}

	return nil
}
