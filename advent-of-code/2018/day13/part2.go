package main

import "fmt"

func part2(puzzleInput []string) any {
	grid, carts := parseInput(puzzleInput)

	for {
		sortCarts(carts)

		for _, cart := range carts {
			if cart.crashed {
				continue
			}

			moveCart(cart)
			handleTrack(cart, grid[cart.y][cart.x])

			if other := findCollision(carts, cart); other != nil {
				cart.crashed = true
				other.crashed = true
			}
		}

		// filter surviving carts
		active := []*Cart{}
		for _, c := range carts {
			if !c.crashed {
				active = append(active, c)
			}
		}
		carts = active

		if len(carts) == 1 {
			return fmt.Sprintf("%d,%d", carts[0].x, carts[0].y)
		}
	}
}
