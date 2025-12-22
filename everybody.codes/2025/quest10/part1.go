package main

func part1(puzzleInput []string) any {
	seen := make(map[position]bool)
	result := 0
	move(puzzleInput, position{i: len(puzzleInput) / 2, j: len(puzzleInput[0]) / 2}, seen, 0, 4)

	for i := 0; i < len(puzzleInput); i++ {
		for j := 0; j < len(puzzleInput[0]); j++ {
			if seen[position{i, j}] && puzzleInput[i][j] == 'S' {
				result++
			}
		}
	}
	return result
}

func move(board []string, p position, seen map[position]bool, d int, md int) int {
	if !isValid(p, len(board), len(board[0])) || d > md {
		return 0
	}
	seen[p] = true
	s := 0
	if board[p.i][p.j] == 'S' {
		s = 1
	}
	for _, m := range moves {
		s += move(board, np(p, m), seen, d+1, md)
	}
	return s
}
