from helpers import apply_phase


def part1(puzzle_input):
    signal = [int(c) for c in puzzle_input[0].strip()]

    for _ in range(100):
        signal = apply_phase(signal)

    return "".join(map(str, signal[:8]))
