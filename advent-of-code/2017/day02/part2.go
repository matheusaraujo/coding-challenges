package main

func part2(puzzleInput []string) interface{} {
	numbers := parseInput(puzzleInput)
	checksum := 0
	for _, row := range numbers {
		for i := 0; i < len(row); i++ {
			for j := i + 1; j < len(row); j++ {
				if row[j]%row[i] == 0 {
					checksum += row[j] / row[i]
					break
				}
			}
		}
	}
	return checksum
}
