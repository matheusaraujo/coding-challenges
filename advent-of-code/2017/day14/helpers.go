package main

import (
	"fmt"
	"strings"
)

func knotHash(input string) string {
	lengths := []int{}
	for _, c := range input {
		lengths = append(lengths, int(c))
	}
	lengths = append(lengths, 17, 31, 73, 47, 23) // from 2017-day10

	size := 256
	list := make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = i
	}

	pos, skip := 0, 0
	for round := 0; round < 64; round++ {
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

	hash := ""
	for i := 0; i < 16; i++ {
		x := list[i*16]
		for j := 1; j < 16; j++ {
			x ^= list[i*16+j]
		}
		hash += fmt.Sprintf("%02x", x)
	}
	return hash
}

func hexToBin(hex string) string {
	bin := ""
	for _, c := range hex {
		var n int
		if c >= '0' && c <= '9' {
			n = int(c - '0')
		} else {
			n = int(c-'a') + 10
		}
		bin += fmt.Sprintf("%04b", n)
	}
	return bin
}

func getRow(key string, i int) string {
	return fmt.Sprintf("%s-%d", key, i)
}

func ones(row string) int {
	return strings.Count(row, "1")
}
