package main

import (
	"sort"
	"strconv"
	"strings"
)

type Interval struct {
	Start, End int
}

func (i Interval) Length() int {
	return i.End - i.Start + 1
}

func parseInput(puzzleInput []string) ([]Interval, []int) {
	intervals := make([]Interval, 0)
	nums := make([]int, 0)
	parsingIntervals := true

	for _, line := range puzzleInput {
		if line == "" {
			parsingIntervals = false
		} else if parsingIntervals {
			interval := strings.Split(line, "-")
			start, _ := strconv.Atoi(interval[0])
			end, _ := strconv.Atoi(interval[1])
			intervals = append(intervals, Interval{start, end})
		} else {
			num, _ := strconv.Atoi(line)
			nums = append(nums, num)
		}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	return intervals, nums
}
