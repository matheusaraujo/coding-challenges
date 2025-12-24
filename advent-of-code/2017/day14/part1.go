package main

func part1(puzzleInput []string) any {
	key := puzzleInput[0]
	used := 0

	for i := 0; i < 128; i++ {
		row := getRow(key, i)
		hash := knotHash(row)
		binRow := hexToBin(hash)
		used += ones(binRow)
	}

	return used
}
