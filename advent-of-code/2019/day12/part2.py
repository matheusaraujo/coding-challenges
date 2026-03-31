from helpers import get_lcm_list, parse_moons


def part2(puzzle_input):
    initial_moons = parse_moons(puzzle_input)
    cycles = [0, 0, 0]

    for d in range(3):
        # Initial state for this dimension
        pos = [m["pos"][d] for m in initial_moons]
        vel = [0] * len(pos)
        initial_state = tuple(zip(pos, vel))

        step = 0
        while True:
            step += 1
            # Gravity: We still need i to slice the inner loop (fixes C0200)
            for i, pos_i in enumerate(pos):
                for j in range(i + 1, len(pos)):
                    if pos_i < pos[j]:
                        vel[i] += 1
                        vel[j] -= 1
                    elif pos_i > pos[j]:
                        vel[i] -= 1
                        vel[j] += 1

            # Velocity: Use enumerate to update in-place (fixes C0200)
            for i, v_val in enumerate(vel):
                pos[i] += v_val

            if tuple(zip(pos, vel)) == initial_state:
                cycles[d] = step
                break

    return get_lcm_list(cycles)
