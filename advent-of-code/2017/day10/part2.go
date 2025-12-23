package main

func part2(puzzleInput []string) any {
	lengths := asciiLengths(puzzleInput[0])
	sparse := runKnot(lengths, 64)
	dense := denseHash(sparse)
	return toHex(dense)
}
