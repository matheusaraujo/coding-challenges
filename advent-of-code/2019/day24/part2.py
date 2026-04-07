# pylint: disable=too-many-branches
def count_neighbors_recursive(levels, d, x, y):
    count = 0
    # Pre-fetch grids to avoid KeyErrors
    current_grid = levels.get(d, "." * 25)
    inner_grid = levels.get(d + 1, "." * 25)
    outer_grid = levels.get(d - 1, "." * 25)

    for dx, dy in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
        nx, ny = x + dx, y + dy

        # 1. Neighbor is OUTSIDE (Level d-1)
        if nx < 0:
            if outer_grid[2 * 5 + 1] == "#":
                count += 1
        elif nx > 4:
            if outer_grid[2 * 5 + 3] == "#":
                count += 1
        elif ny < 0:
            if outer_grid[1 * 5 + 2] == "#":
                count += 1
        elif ny > 4:
            if outer_grid[3 * 5 + 2] == "#":
                count += 1

        # 2. Neighbor is the CENTER (Level d+1)
        elif nx == 2 and ny == 2:
            if x == 1:  # Moving Right into center: check left edge of inner
                count += sum(1 for r in range(5) if inner_grid[r * 5 + 0] == "#")
            elif x == 3:  # Moving Left into center: check right edge of inner
                count += sum(1 for r in range(5) if inner_grid[r * 5 + 4] == "#")
            elif y == 1:  # Moving Down into center: check top edge of inner
                count += sum(1 for c in range(5) if inner_grid[0 * 5 + c] == "#")
            elif y == 3:  # Moving Up into center: check bottom edge of inner
                count += sum(1 for c in range(5) if inner_grid[4 * 5 + c] == "#")

        # 3. Neighbor is a normal tile on the SAME level
        else:
            if current_grid[ny * 5 + nx] == "#":
                count += 1
    return count


# pylint: disable=too-many-locals
def part2(puzzle_input, minutes=200):
    if isinstance(puzzle_input, str):
        lines = [
            line.strip() for line in puzzle_input.strip().split("\n") if line.strip()
        ]
    else:
        lines = [line.strip() for line in puzzle_input if line.strip()]

    initial_grid = "".join(lines)
    levels = {0: initial_grid}

    for _ in range(minutes):
        new_levels = {}
        # The range of levels expands by 1 in each direction every minute
        min_d = min(levels.keys()) - 1
        max_d = max(levels.keys()) + 1

        for d in range(min_d, max_d + 1):
            grid = levels.get(d, "." * 25)
            new_grid_chars = list("." * 25)

            for i in range(25):
                if i == 12:
                    continue  # Center tile is the portal to next level
                x, y = i % 5, i // 5
                neighbors = count_neighbors_recursive(levels, d, x, y)

                if grid[i] == "#" and neighbors == 1:
                    new_grid_chars[i] = "#"
                elif grid[i] == "." and neighbors in [1, 2]:
                    new_grid_chars[i] = "#"

            res = "".join(new_grid_chars)
            if "#" in res:
                new_levels[d] = res
        levels = new_levels

    return sum(grid.count("#") for grid in levels.values())
