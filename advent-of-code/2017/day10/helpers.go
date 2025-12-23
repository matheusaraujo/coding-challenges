package main

import (
	"fmt"
	"strconv"
	"strings"
)

func runKnot(lengths []int, rounds int) []int {
	const size = 256

	list := make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = i
	}

	pos := 0
	skip := 0

	for r := 0; r < rounds; r++ {
		for _, length := range lengths {
			for i := 0; i < length/2; i++ {
				a := (pos + i) % size
				b := (pos + length - 1 - i) % size
				list[a], list[b] = list[b], list[a]
			}
			pos = (pos + length + skip) % size
			skip++
		}
	}

	return list
}

func parseCommaInts(s string) []int {
	parts := strings.Split(strings.TrimSpace(s), ",")
	out := make([]int, len(parts))
	for i, p := range parts {
		out[i], _ = strconv.Atoi(p)
	}
	return out
}

func asciiLengths(s string) []int {
	out := make([]int, 0, len(s)+5)
	for _, b := range []byte(strings.TrimSpace(s)) {
		out = append(out, int(b))
	}
	return append(out, 17, 31, 73, 47, 23)
}

func denseHash(sparse []int) []int {
	dense := make([]int, 16)
	for i := 0; i < 16; i++ {
		x := sparse[i*16]
		for j := 1; j < 16; j++ {
			x ^= sparse[i*16+j]
		}
		dense[i] = x
	}
	return dense
}

func toHex(dense []int) string {
	var sb strings.Builder
	for _, v := range dense {
		sb.WriteString(fmt.Sprintf("%02x", v))
	}
	return sb.String()
}
