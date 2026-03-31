package main

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Group struct {
	id         int
	army       string
	units      int
	hp         int
	attack     int
	attackType string
	initiative int
	weak       map[string]bool
	immune     map[string]bool
}

func (g *Group) effectivePower() int {
	return g.units * g.attack
}

func (g *Group) damageTo(def *Group) int {
	if def.immune[g.attackType] {
		return 0
	}
	dmg := g.effectivePower()
	if def.weak[g.attackType] {
		dmg *= 2
	}
	return dmg
}

var lineRegex = regexp.MustCompile(`(\d+) units each with (\d+) hit points (?:\((.*?)\) )?with an attack that does (\d+) (\w+) damage at initiative (\d+)`)

func parseInput(input []string, boost int) []*Group {
	var groups []*Group
	army := ""
	id := 0

	for _, line := range input {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if line == "Immune System:" {
			army = "Immune"
			continue
		}
		if line == "Infection:" {
			army = "Infection"
			continue
		}

		m := lineRegex.FindStringSubmatch(line)

		units, _ := strconv.Atoi(m[1])
		hp, _ := strconv.Atoi(m[2])
		attack, _ := strconv.Atoi(m[4])
		attackType := m[5]
		initiative, _ := strconv.Atoi(m[6])

		if army == "Immune" {
			attack += boost
		}

		weak := map[string]bool{}
		immune := map[string]bool{}

		if m[3] != "" {
			parts := strings.Split(m[3], ";")
			for _, p := range parts {
				p = strings.TrimSpace(p)
				if strings.HasPrefix(p, "weak to ") {
					for _, t := range strings.Split(p[8:], ", ") {
						weak[t] = true
					}
				}
				if strings.HasPrefix(p, "immune to ") {
					for _, t := range strings.Split(p[10:], ", ") {
						immune[t] = true
					}
				}
			}
		}

		groups = append(groups, &Group{
			id:         id,
			army:       army,
			units:      units,
			hp:         hp,
			attack:     attack,
			attackType: attackType,
			initiative: initiative,
			weak:       weak,
			immune:     immune,
		})
		id++
	}

	return groups
}

func simulate(input []string, boost int) (string, int, bool) {
	groups := parseInput(input, boost)

	for {
		// target selection
		sort.Slice(groups, func(i, j int) bool {
			if groups[i].effectivePower() == groups[j].effectivePower() {
				return groups[i].initiative > groups[j].initiative
			}
			return groups[i].effectivePower() > groups[j].effectivePower()
		})

		targets := map[*Group]*Group{}
		chosen := map[*Group]bool{}

		for _, g := range groups {
			var best *Group
			bestDmg := 0

			for _, e := range groups {
				if e.army == g.army || chosen[e] || e.units <= 0 {
					continue
				}
				dmg := g.damageTo(e)
				if dmg == 0 {
					continue
				}
				if dmg > bestDmg ||
					(dmg == bestDmg && (e.effectivePower() > best.effectivePower() ||
						(e.effectivePower() == best.effectivePower() && e.initiative > best.initiative))) {
					best = e
					bestDmg = dmg
				}
			}

			if best != nil {
				targets[g] = best
				chosen[best] = true
			}
		}

		// attack phase
		sort.Slice(groups, func(i, j int) bool {
			return groups[i].initiative > groups[j].initiative
		})

		totalKilled := 0

		for _, g := range groups {
			if g.units <= 0 {
				continue
			}
			target := targets[g]
			if target == nil {
				continue
			}

			damage := g.damageTo(target)
			killed := damage / target.hp
			if killed > target.units {
				killed = target.units
			}

			target.units -= killed
			totalKilled += killed
		}

		if totalKilled == 0 {
			return "", 0, true // stalemate
		}

		// remove dead
		var alive []*Group
		for _, g := range groups {
			if g.units > 0 {
				alive = append(alive, g)
			}
		}
		groups = alive

		// check winner
		army := groups[0].army
		allSame := true
		totalUnits := 0

		for _, g := range groups {
			totalUnits += g.units
			if g.army != army {
				allSame = false
			}
		}

		if allSame {
			return army, totalUnits, false
		}
	}
}
