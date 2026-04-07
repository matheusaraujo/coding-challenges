def get_biodiversity(grid):
    score = 0
    for i, char in enumerate(grid):
        if char == "#":
            score += 2 ** i
    return score


def parse_grid(puzzle_input):
    if isinstance(puzzle_input, str):
        lines = puzzle_input.strip().split("\n")
    else:
        lines = [line.strip() for line in puzzle_input if line.strip()]
    return "".join(lines)


def simulate_step(grid):
    new_grid = list(grid)
    for i in range(25):
        x, y = i % 5, i // 5
        neighbors = 0
        for dx, dy in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
            nx, ny = x + dx, y + dy
            if 0 <= nx < 5 and 0 <= ny < 5:
                if grid[ny * 5 + nx] == "#":
                    neighbors += 1

        if grid[i] == "#" and neighbors != 1:
            new_grid[i] = "."
        elif grid[i] == "." and neighbors in [1, 2]:
            new_grid[i] = "#"
    return "".join(new_grid)
