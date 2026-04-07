from collections import deque

from helpers import is_outer, parse_maze


# pylint: disable=too-many-locals
def part2(puzzle_input):
    grid, portals, label_coords, max_x, max_y = parse_maze(puzzle_input)

    start = label_coords["AA"][0]
    goal = label_coords["ZZ"][0]

    # State: (x, y, level)
    queue = deque([(start[0], start[1], 0, 0)])  # x, y, level, dist
    visited = {(start[0], start[1], 0)}

    while queue:
        x, y, level, dist = queue.popleft()

        if (x, y) == goal and level == 0:
            return dist

        # 1. Normal Walking
        for dx, dy in [(0, 1), (1, 0), (0, -1), (-1, 0)]:
            nx, ny = x + dx, y + dy
            if grid.get((nx, ny)) == "." and (nx, ny, level) not in visited:
                visited.add((nx, ny, level))
                queue.append((nx, ny, level, dist + 1))

        # 2. Recursive Warping
        if (x, y) in portals:
            label = portals[(x, y)]
            if label in ("AA", "ZZ"):
                continue

            # Determine if we are going deeper (+1) or shallower (-1)
            if is_outer((x, y), max_x, max_y):
                if level == 0:
                    continue  # Outer portals at level 0 are walls
                new_level = level - 1
            else:
                new_level = level + 1

            others = label_coords[label]
            dest = others[0] if others[1] == (x, y) else others[1]

            if (dest[0], dest[1], new_level) not in visited:
                visited.add((dest[0], dest[1], new_level))
                queue.append((dest[0], dest[1], new_level, dist + 1))

    return None
