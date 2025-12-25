package main

func part2(puzzleInput []string) any {
	particles := parseInput(puzzleInput)

	const ticks = 2000

	for t := 0; t < ticks; t++ {
		positions := make(map[[3]int][]int)

		for i := range particles {
			if !particles[i].alive {
				continue
			}
			step(&particles[i])
			positions[particles[i].p] = append(positions[particles[i].p], i)
		}

		for _, idxs := range positions {
			if len(idxs) > 1 {
				for _, i := range idxs {
					particles[i].alive = false
				}
			}
		}
	}

	count := 0
	for _, p := range particles {
		if p.alive {
			count++
		}
	}

	return count
}

func step(p *particle) {
	for d := 0; d < 3; d++ {
		p.v[d] += p.a[d]
		p.p[d] += p.v[d]
	}
}
