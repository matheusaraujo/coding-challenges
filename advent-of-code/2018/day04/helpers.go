package main

import (
	"sort"
	"strconv"
	"strings"
)

func parseInput(puzzleInput []string) map[int][60]int {
	sort.Strings(puzzleInput)

	sleepMap := make(map[int][60]int)

	var currentGuard int
	var sleepStart int

	for _, line := range puzzleInput {
		if strings.Contains(line, "Guard") {
			parts := strings.Split(line, " ")
			idStr := strings.TrimPrefix(parts[3], "#")
			id, _ := strconv.Atoi(idStr)
			currentGuard = id
		} else if strings.Contains(line, "falls asleep") {
			minute, _ := strconv.Atoi(line[15:17])
			sleepStart = minute
		} else if strings.Contains(line, "wakes up") {
			wakeMinute, _ := strconv.Atoi(line[15:17])

			minutes := sleepMap[currentGuard]
			for m := sleepStart; m < wakeMinute; m++ {
				minutes[m]++
			}
			sleepMap[currentGuard] = minutes
		}
	}

	return sleepMap
}
