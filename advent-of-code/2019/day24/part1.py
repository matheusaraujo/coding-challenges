from helpers import get_biodiversity, parse_grid, simulate_step


def part1(puzzle_input):
    grid = parse_grid(puzzle_input)
    seen = {grid}

    while True:
        grid = simulate_step(grid)
        if grid in seen:
            return get_biodiversity(grid)
        seen.add(grid)
