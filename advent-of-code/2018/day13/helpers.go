package main

import (
	"sort"
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Turn int

const (
	TurnLeft Turn = iota
	Straight
	TurnRight
)

type Cart struct {
	x, y     int
	dir      Direction
	nextTurn Turn
	crashed  bool
}

func parseInput(input []string) ([][]rune, []*Cart) {
	grid := make([][]rune, len(input))
	var carts []*Cart

	for y, line := range input {
		row := []rune(line)
		for x, ch := range row {
			switch ch {
			case '^':
				carts = append(carts, &Cart{x, y, Up, TurnLeft, false})
				row[x] = '|'
			case 'v':
				carts = append(carts, &Cart{x, y, Down, TurnLeft, false})
				row[x] = '|'
			case '<':
				carts = append(carts, &Cart{x, y, Left, TurnLeft, false})
				row[x] = '-'
			case '>':
				carts = append(carts, &Cart{x, y, Right, TurnLeft, false})
				row[x] = '-'
			}
		}
		grid[y] = row
	}

	return grid, carts
}

func moveCart(c *Cart) {
	switch c.dir {
	case Up:
		c.y--
	case Down:
		c.y++
	case Left:
		c.x--
	case Right:
		c.x++
	}
}

func turnLeft(d Direction) Direction {
	return (d + 3) % 4
}

func turnRight(d Direction) Direction {
	return (d + 1) % 4
}

func handleTrack(c *Cart, track rune) {
	switch track {
	case '/':
		switch c.dir {
		case Up, Down:
			c.dir = turnRight(c.dir)
		case Left, Right:
			c.dir = turnLeft(c.dir)
		}
	case '\\':
		switch c.dir {
		case Up, Down:
			c.dir = turnLeft(c.dir)
		case Left, Right:
			c.dir = turnRight(c.dir)
		}
	case '+':
		switch c.nextTurn {
		case TurnLeft:
			c.dir = turnLeft(c.dir)
		case Straight:
			// no change
		case TurnRight:
			c.dir = turnRight(c.dir)
		}
		c.nextTurn = (c.nextTurn + 1) % 3
	}
}

func sortCarts(carts []*Cart) {
	sort.Slice(carts, func(i, j int) bool {
		if carts[i].y == carts[j].y {
			return carts[i].x < carts[j].x
		}
		return carts[i].y < carts[j].y
	})
}

func findCollision(carts []*Cart, current *Cart) *Cart {
	for _, other := range carts {
		if other != current && !other.crashed &&
			other.x == current.x && other.y == current.y {
			return other
		}
	}
	return nil
}
