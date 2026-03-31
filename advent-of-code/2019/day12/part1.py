from helpers import parse_moons


def part1(puzzle_input):
    moons = parse_moons(puzzle_input)

    for _ in range(1000):
        # Apply Gravity
        for i, moon_a in enumerate(moons):
            for moon_b in moons[i + 1 :]:
                for d in range(3):
                    if moon_a["pos"][d] < moon_b["pos"][d]:
                        moon_a["vel"][d] += 1
                        moon_b["vel"][d] -= 1
                    elif moon_a["pos"][d] > moon_b["pos"][d]:
                        moon_a["vel"][d] -= 1
                        moon_b["vel"][d] += 1

        # Apply Velocity
        for m in moons:
            for d in range(3):
                m["pos"][d] += m["vel"][d]

    total_energy = 0
    for m in moons:
        pot = sum(abs(x) for x in m["pos"])
        kin = sum(abs(x) for x in m["vel"])
        total_energy += pot * kin

    return total_energy
