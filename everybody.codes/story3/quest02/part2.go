package main

func part2(puzzleInput []string) any {

	start, bones := ParseGrid(puzzleInput)

	var bone Pos
	for b := range bones {
		bone = b
		break
	}

	visited := map[Pos]bool{}
	pos := start
	visited[pos] = true

	dir := 0
	steps := 0

	for {

		for {
			next := Add(pos, Cardinal[dir])

			if !visited[next] && next != bone {

				pos = next
				visited[pos] = true
				steps++
				dir = (dir + 1) % 4
				break
			}

			dir = (dir + 1) % 4
		}

		surrounded := true

		for _, d := range Cardinal {
			n := Add(bone, d)
			if !visited[n] {
				surrounded = false
				break
			}
		}

		if surrounded {
			return steps
		}
	}
}
