package main

func part2(puzzleInput []string) any {
	samples, program := parseInput(puzzleInput)

	possibleOps := make(map[int]map[string]bool)
	for i := 0; i < 16; i++ {
		possibleOps[i] = make(map[string]bool)
		for name := range Opcodes {
			possibleOps[i][name] = true
		}
	}

	for _, s := range samples {
		opNum := s.Inst[0]
		for name, opFunc := range Opcodes {
			if opFunc(s.Before, s.Inst) != s.After {
				delete(possibleOps[opNum], name)
			}
		}
	}

	finalMapping := make(map[int]string)
	for len(finalMapping) < 16 {
		for num, ops := range possibleOps {
			if len(ops) == 1 {
				var foundName string
				for name := range ops {
					foundName = name
				}
				finalMapping[num] = foundName

				for otherNum := range possibleOps {
					delete(possibleOps[otherNum], foundName)
				}
			}
		}
	}

	registers := [4]int{0, 0, 0, 0}
	for _, inst := range program {
		opName := finalMapping[inst[0]]
		registers = Opcodes[opName](registers, inst)
	}

	return registers[0]
}
