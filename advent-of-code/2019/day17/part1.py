# part1.py

from helpers import IntcodeComputer, get_grid, parse_program


def part1(puzzle_input):
    program = parse_program(puzzle_input)
    computer = IntcodeComputer(program)
    output = computer.run()

    grid = get_grid(output)

    h, w = len(grid), len(grid[0])
    total = 0

    for y in range(1, h - 1):
        for x in range(1, w - 1):
            if (
                grid[y][x] == "#"
                and grid[y - 1][x] == "#"
                and grid[y + 1][x] == "#"
                and grid[y][x - 1] == "#"
                and grid[y][x + 1] == "#"
            ):
                total += x * y

    return total
