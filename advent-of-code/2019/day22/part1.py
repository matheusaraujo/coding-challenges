from helpers import parse_instructions


def part1(puzzle_input):
    n = 10007
    target_card = 2019
    a, b = parse_instructions(puzzle_input, n)

    # Result is f(2019) = a*2019 + b % n
    return (a * target_card + b) % n
