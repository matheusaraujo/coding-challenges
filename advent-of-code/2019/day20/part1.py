from collections import deque

from helpers import parse_maze


# pylint: disable=too-many-locals
def part1(puzzle_input):
    grid, portals, label_coords, _, _ = parse_maze(puzzle_input)

    start = label_coords["AA"][0]
    goal = label_coords["ZZ"][0]

    queue = deque([(start, 0)])
    visited = {start}

    while queue:
        (x, y), dist = queue.popleft()
        if (x, y) == goal:
            return dist

        # Standard moves
        neighbors = [(x + 1, y), (x - 1, y), (x, y + 1), (x, y - 1)]

        # Warp moves
        if (x, y) in portals:
            label = portals[(x, y)]
            if label not in ("AA", "ZZ"):
                # Find the other end of the portal
                others = label_coords[label]
                dest = others[0] if others[1] == (x, y) else others[1]
                neighbors.append(dest)

        for nx, ny in neighbors:
            if grid.get((nx, ny)) == "." and (nx, ny) not in visited:
                visited.add((nx, ny))
                queue.append(((nx, ny), dist + 1))
    return None
