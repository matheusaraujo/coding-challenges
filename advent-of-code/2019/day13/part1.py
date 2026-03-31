from helpers import Intcode


def part1(puzzle_input):
    raw_data = "".join(puzzle_input) if isinstance(puzzle_input, list) else puzzle_input
    program = [int(x) for x in raw_data.strip().split(",")]

    computer = Intcode(program)
    grid = {}

    while not computer.halted:
        x = computer.run([])
        y = computer.run([])
        tile_id = computer.run([])

        if x is None:
            break
        grid[(x, y)] = tile_id

    return sum(1 for tile in grid.values() if tile == 2)
