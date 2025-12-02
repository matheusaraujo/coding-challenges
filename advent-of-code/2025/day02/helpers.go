package main

import (
	"strconv"
	"strings"
)

func solve(puzzleInput []string, check func(string) bool) interface{} {
	pairs := strings.Split(puzzleInput[0], ",")
	result := 0
	for _, pair := range pairs {
		numbers := strings.Split(pair, "-")
		j, _ := strconv.Atoi(numbers[0])
		k, _ := strconv.Atoi(numbers[1])
		for i := j; i <= k; i++ {
			s := strconv.Itoa(i)
			if check(s) {
				result += i
			}
		}
	}
	return result
}
