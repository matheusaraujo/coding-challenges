package main

import (
	"strconv"
)

func part2(puzzleInput []string) string {
	intervals, _ := parseInput(puzzleInput)
	return strconv.Itoa(merge(intervals))
}

func merge(intervals []Interval) int {
	merged := []Interval{intervals[0]}
	for _, curr := range intervals[1:] {
		last := &merged[len(merged)-1]
		if curr.Start <= last.End {
			if curr.End > last.End {
				last.End = curr.End
			}
		} else {
			merged = append(merged, curr)
		}
	}
	length := 0
	for _, i := range merged {
		length += i.Length()
	}

	return length
}
