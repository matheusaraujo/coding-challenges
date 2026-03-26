// based on https://www.reddit.com/r/adventofcode/comments/a6chwa/comment/ebtyelm/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button

package main

import (
	"sort"
	"strings"
)

type Point struct {
	X, Y int
}

type Unit struct {
	unitType     string
	x, y         int
	isAlive      bool
	hp           int
	attackDamage int
}

func (u *Unit) Pos() Point {
	return Point{u.x, u.y}
}

func (u *Unit) Attack(damage int) {
	if u.isAlive {
		u.hp -= damage
		if u.hp <= 0 {
			u.isAlive = false
		}
	}
}

func neighbours(x, y int) []Point {
	return []Point{
		{x - 1, y},
		{x + 1, y},
		{x, y - 1},
		{x, y + 1},
	}
}

func readingOrder(p Point) (int, int) {
	return p.Y, p.X
}

// BFS to find closest targets
func findClosest(graph map[Point][]Point, excluded map[Point]bool, start Point, targets map[Point]bool) ([]Point, *int) {
	if _, ok := graph[start]; !ok {
		return nil, nil
	}

	seen := map[Point]bool{}
	queue := []struct {
		p    Point
		dist int
	}{{start, 0}}

	var foundDist *int
	var closest []Point

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if foundDist != nil && cur.dist > *foundDist {
			return closest, foundDist
		}

		if seen[cur.p] || excluded[cur.p] {
			continue
		}

		seen[cur.p] = true

		if targets[cur.p] {
			d := cur.dist
			foundDist = &d
			closest = append(closest, cur.p)
		}

		for _, n := range graph[cur.p] {
			if !seen[n] {
				queue = append(queue, struct {
					p    Point
					dist int
				}{n, cur.dist + 1})
			}
		}
	}

	return closest, foundDist
}

func solve(lines []string, d int, noElvesDie bool) (bool, int) {
	data := [][]string{}
	for _, line := range lines {
		row := strings.Split(line, "")
		data = append(data, row)
	}

	var units []*Unit

	// Extract units
	for y := range data {
		for x := range data[0] {
			t := data[y][x]
			if t == "G" || t == "E" {
				damage := 3
				if t == "E" {
					damage = d
				}
				units = append(units, &Unit{t, x, y, true, 200, damage})
				data[y][x] = "."
			}
		}
	}

	// Build graph
	graph := map[Point][]Point{}
	for y := range data {
		for x := range data[0] {
			if data[y][x] == "." {
				p := Point{x, y}
				for _, n := range neighbours(x, y) {
					if n.X >= 0 && n.X < len(data[0]) && n.Y >= 0 && n.Y < len(data) {
						if data[n.Y][n.X] == "." {
							graph[p] = append(graph[p], n)
						}
					}
				}
			}
		}
	}

	round := 0

	for {
		sort.Slice(units, func(i, j int) bool {
			yi, xi := readingOrder(units[i].Pos())
			yj, xj := readingOrder(units[j].Pos())
			if yi == yj {
				return xi < xj
			}
			return yi < yj
		})

		for idx, c := range units {
			if !c.isAlive {
				continue
			}

			// Find enemies
			var enemies []*Unit
			enemyPositions := map[Point]bool{}

			for _, u := range units {
				if u.isAlive && u.unitType != c.unitType {
					enemies = append(enemies, u)
					enemyPositions[u.Pos()] = true
				}
			}

			nearby := neighbours(c.x, c.y)

			var inRange []Point
			for _, p := range nearby {
				if enemyPositions[p] {
					inRange = append(inRange, p)
				}
			}

			// Move if needed
			if len(inRange) == 0 {
				surrounding := map[Point]bool{}
				for _, e := range enemies {
					for _, n := range neighbours(e.x, e.y) {
						if _, ok := graph[n]; ok {
							surrounding[n] = true
						}
					}
				}

				excluded := map[Point]bool{}
				for _, u := range units {
					if u.isAlive && u != c {
						excluded[u.Pos()] = true
					}
				}

				closest, dist := findClosest(graph, excluded, c.Pos(), surrounding)

				if dist != nil && len(closest) > 0 {
					sort.Slice(closest, func(i, j int) bool {
						yi, xi := readingOrder(closest[i])
						yj, xj := readingOrder(closest[j])
						if yi == yj {
							return xi < xj
						}
						return yi < yj
					})

					choice := closest[0]

					sort.Slice(nearby, func(i, j int) bool {
						yi, xi := readingOrder(nearby[i])
						yj, xj := readingOrder(nearby[j])
						if yi == yj {
							return xi < xj
						}
						return yi < yj
					})

					for _, s := range nearby {
						_, d2 := findClosest(graph, excluded, s, map[Point]bool{choice: true})
						if d2 != nil && *d2 == *dist-1 {
							c.x = s.X
							c.y = s.Y
							break
						}
					}
				}

				// recompute in range
				inRange = nil
				for _, p := range neighbours(c.x, c.y) {
					if enemyPositions[p] {
						inRange = append(inRange, p)
					}
				}
			}

			// Attack
			if len(inRange) > 0 {
				var targets []*Unit
				for _, e := range enemies {
					if containsPoint(inRange, e.Pos()) {
						targets = append(targets, e)
					}
				}

				sort.Slice(targets, func(i, j int) bool {
					if targets[i].hp == targets[j].hp {
						yi, xi := readingOrder(targets[i].Pos())
						yj, xj := readingOrder(targets[j].Pos())
						if yi == yj {
							return xi < xj
						}
						return yi < yj
					}
					return targets[i].hp < targets[j].hp
				})

				target := targets[0]
				target.Attack(c.attackDamage)

				if noElvesDie && target.unitType == "E" && !target.isAlive {
					return false, 0
				}

				aliveTypes := map[string]bool{}
				sumHP := 0

				for _, u := range units {
					if u.isAlive {
						aliveTypes[u.unitType] = true
						sumHP += u.hp
					}
				}

				if len(aliveTypes) == 1 {
					if idx == len(units)-1 {
						round++
					}
					return true, round * sumHP
				}
			}
		}
		round++
	}
}

func containsPoint(list []Point, p Point) bool {
	for _, x := range list {
		if x == p {
			return true
		}
	}
	return false
}
