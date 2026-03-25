package main

import "fmt"

func part1(puzzleInput []string) any {
	grid, carts := parseInput(puzzleInput)

	for {
		sortCarts(carts)

		for _, cart := range carts {
			moveCart(cart)
			handleTrack(cart, grid[cart.y][cart.x])

			if other := findCollision(carts, cart); other != nil {
				return fmt.Sprintf("%d,%d", cart.x, cart.y)
			}
		}
	}
}
