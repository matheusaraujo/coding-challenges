package main

import (
	"math"
	"strconv"
	"strings"
)

// based on https://www.reddit.com/r/adventofcode/comments/7lms6p/comment/drnh5sx/
func part2(puzzleInput []string) any {
	h := 0
	x, _ := strconv.Atoi(strings.Fields(puzzleInput[0])[2])

	b := x*100 + 100000
	c := b + 17000

	for n := b; n <= c; n += 17 {
		if !isPrime(n) {
			h++
		}
	}

	return h
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrtN; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
