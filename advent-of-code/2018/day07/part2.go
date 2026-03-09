package main

import "sort"

type Worker struct {
	task      byte
	remaining int
	busy      bool
}

func part2(puzzleInput []string) any {
	nodes, graph, inDegree := parseSteps(puzzleInput)

	const workerCount = 5

	workers := make([]Worker, workerCount)
	done := map[byte]bool{}
	inProgress := map[byte]bool{}

	time := 0

	for len(done) < len(nodes) {

		// find available tasks
		available := []byte{}
		for _, n := range nodes {
			if !done[n] && !inProgress[n] && inDegree[n] == 0 {
				available = append(available, n)
			}
		}

		sort.Slice(available, func(i, j int) bool {
			return available[i] < available[j]
		})

		// assign tasks
		for _, step := range available {
			for i := range workers {
				if !workers[i].busy {
					workers[i] = Worker{
						task:      step,
						remaining: stepDuration(step),
						busy:      true,
					}
					inProgress[step] = true
					break
				}
			}
		}

		// advance time
		time++

		for i := range workers {
			if !workers[i].busy {
				continue
			}

			workers[i].remaining--

			if workers[i].remaining == 0 {
				step := workers[i].task

				done[step] = true
				delete(inProgress, step)

				for _, next := range graph[step] {
					inDegree[next]--
				}

				workers[i].busy = false
			}
		}
	}

	return time
}

func stepDuration(step byte) int {
	return 60 + int(step-'A'+1)
}
