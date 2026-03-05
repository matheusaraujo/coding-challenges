package main

func part3(puzzleInput []string) any {

	groups := map[string]int{
		"red-matte":   0,
		"red-shiny":   0,
		"green-matte": 0,
		"green-shiny": 0,
		"blue-matte":  0,
		"blue-shiny":  0,
	}

	for _, line := range puzzleInput {
		c := NewComponent(line)

		shineType := ""
		if c.Ns <= 30 {
			shineType = "matte"
		} else if c.Ns >= 33 {
			shineType = "shiny"
		} else {
			continue
		}

		color := ""
		if c.Nr > c.Ng && c.Nr > c.Nb {
			color = "red"
		} else if c.Ng > c.Nr && c.Ng > c.Nb {
			color = "green"
		} else if c.Nb > c.Nr && c.Nb > c.Ng {
			color = "blue"
		} else {
			continue
		}

		group := color + "-" + shineType
		groups[group] += c.Id
	}

	max := 0
	for _, v := range groups {
		if v > max {
			max = v
		}
	}

	return max
}
