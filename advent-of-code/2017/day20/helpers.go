package main

import (
	"regexp"
	"strconv"
)

type particle struct {
	id    int
	p     [3]int
	v     [3]int
	a     [3]int
	alive bool
}

func parseInput(input []string) []particle {
	re := regexp.MustCompile(`-?\d+`)
	particles := make([]particle, len(input))

	for i, line := range input {
		numsStr := re.FindAllString(line, -1)
		nums := make([]int, len(numsStr))
		for j, s := range numsStr {
			nums[j], _ = strconv.Atoi(s)
		}

		particles[i] = particle{
			id:    i,
			p:     [3]int{nums[0], nums[1], nums[2]},
			v:     [3]int{nums[3], nums[4], nums[5]},
			a:     [3]int{nums[6], nums[7], nums[8]},
			alive: true,
		}
	}

	return particles
}
