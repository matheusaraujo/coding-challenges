package main

func step(recipes *[]int, e1, e2 *int) {
	sum := (*recipes)[*e1] + (*recipes)[*e2]

	// append digits of sum
	if sum >= 10 {
		*recipes = append(*recipes, sum/10)
	}
	*recipes = append(*recipes, sum%10)

	// move elves
	*e1 = (*e1 + 1 + (*recipes)[*e1]) % len(*recipes)
	*e2 = (*e2 + 1 + (*recipes)[*e2]) % len(*recipes)
}

func endsWith(recipes []int, target []int) bool {
	if len(recipes) < len(target) {
		return false
	}
	start := len(recipes) - len(target)
	for i := range target {
		if recipes[start+i] != target[i] {
			return false
		}
	}
	return true
}

func endsWithOffset(recipes []int, target []int, offset int) bool {
	if len(recipes) < len(target)+offset {
		return false
	}
	start := len(recipes) - len(target) - offset
	for i := range target {
		if recipes[start+i] != target[i] {
			return false
		}
	}
	return true
}
