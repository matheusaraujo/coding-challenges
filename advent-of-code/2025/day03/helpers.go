package main

import (
	"strconv"
)

func btoi(s string, i int) int {
	r, _ := strconv.Atoi(string(s[i]))
	return r
}

func atoi(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}
