from helpers import check_pos


def part1(puzzle_input):
    if isinstance(puzzle_input, str):
        program = [int(x) for x in puzzle_input.strip().split(",")]
    else:
        program = [int(x) for x in puzzle_input[0].strip().split(",")]

    count = 0
    for y in range(50):
        for x in range(50):
            count += check_pos(program, x, y)
    return count
