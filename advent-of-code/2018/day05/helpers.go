package main

func reacts(a, b byte) bool {
	if a > b {
		return a-b == 32
	}
	return b-a == 32
}

func reactPolymer(polymer []byte) int {
	stack := make([]byte, 0, len(polymer))

	for _, c := range polymer {
		if len(stack) > 0 && reacts(stack[len(stack)-1], c) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}

	return len(stack)
}
