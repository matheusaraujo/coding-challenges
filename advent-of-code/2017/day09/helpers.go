package main

func solve(line string) (int, int) {
	depth, score, chars, garbage, skip := 0, 0, 0, false, false

	for _, c := range line {
		if skip {
			skip = false
			continue
		}

		if garbage {
			if c == '!' {
				skip = true
			} else if c == '>' {
				garbage = false
			} else {
				chars++
			}
		} else {
			if c == '<' {
				garbage = true
			} else if c == '{' {
				depth++
				score += depth
			} else if c == '}' {
				depth--
			}
		}
	}

	return score, chars
}
