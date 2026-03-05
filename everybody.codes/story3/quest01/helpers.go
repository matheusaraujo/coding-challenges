package main

import (
	"strconv"
	"strings"
)

type Component struct {
	Id int

	R string
	G string
	B string
	S string

	Nr int
	Ng int
	Nb int
	Ns int

	Nc int
}

func NewComponent(line string) Component {
	parts := strings.Split(line, ":")

	id, _ := strconv.Atoi(parts[0])
	colors := strings.Split(parts[1], " ")

	c := Component{}
	c.Id = id

	c.R = colors[0]
	c.Nr = parseColor(c.R)

	c.G = colors[1]
	c.Ng = parseColor(c.G)

	c.B = colors[2]
	c.Nb = parseColor(c.B)

	if len(colors) > 3 {
		c.S = colors[3]
		c.Ns = parseColor(c.S)
	}

	c.Nc = c.Nr + c.Ng + c.Nb

	return c
}

func parseColor(color string) int {
	result := 0

	for i := 0; i < len(color); i++ {
		if color[i] >= 'A' && color[i] <= 'Z' {
			result += 1 << (len(color) - i - 1)
		}
	}

	return result
}
