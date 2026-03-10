package main

type Pos struct {
	R int
	C int
}

func Add(a, b Pos) Pos {
	return Pos{a.R + b.R, a.C + b.C}
}

var Cardinal = []Pos{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

var SpiralDirs = []Pos{
	{-1, 0}, {-1, 0}, {-1, 0},
	{0, 1}, {0, 1}, {0, 1},
	{1, 0}, {1, 0}, {1, 0},
	{0, -1}, {0, -1}, {0, -1},
}

func ParseGrid(grid []string) (Pos, map[Pos]bool) {

	start := Pos{}
	bones := map[Pos]bool{}

	for r := range grid {
		for c, ch := range grid[r] {

			p := Pos{r, c}

			if ch == '@' {
				start = p
			}

			if ch == '#' {
				bones[p] = true
			}
		}
	}

	return start, bones
}
