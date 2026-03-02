package main

import (
	"strconv"
	"strings"
)

type claim struct {
	id     int
	left   int
	top    int
	width  int
	height int
}

func parseLine(line string) claim {
	parts := strings.Split(line, " ")

	idStr := strings.TrimPrefix(parts[0], "#")
	id, _ := strconv.Atoi(idStr)

	coords := strings.TrimSuffix(parts[2], ":")
	coordParts := strings.Split(coords, ",")
	left, _ := strconv.Atoi(coordParts[0])
	top, _ := strconv.Atoi(coordParts[1])

	sizeParts := strings.Split(parts[3], "x")
	width, _ := strconv.Atoi(sizeParts[0])
	height, _ := strconv.Atoi(sizeParts[1])

	return claim{
		id:     id,
		left:   left,
		top:    top,
		width:  width,
		height: height,
	}
}
