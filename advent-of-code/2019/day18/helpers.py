from collections import deque


def parse_grid(puzzle_input):
    grid = {}
    starts = []
    keys = {}

    # Handle both raw string input and list of strings
    if isinstance(puzzle_input, str):
        lines = puzzle_input.strip().split("\n")
    else:
        lines = [line.strip() for line in puzzle_input if line.strip()]

    for y, line in enumerate(lines):
        for x, char in enumerate(line):
            grid[(x, y)] = char
            if char == "@":
                starts.append((x, y))
            elif "a" <= char <= "z":
                keys[char] = (x, y)
    return grid, starts, keys


def get_reachable_keys(grid, start_pos, current_keys):
    """BFS to find all keys reachable from a position given current keys."""
    queue = deque([(start_pos, 0)])
    visited = {start_pos}
    reachable = {}

    while queue:
        (x, y), dist = queue.popleft()
        char = grid.get((x, y), "#")

        # If we hit a key we don't have, it's a potential destination
        if "a" <= char <= "z" and char not in current_keys:
            reachable[char] = (dist, (x, y))
            continue  # Stop at the key; don't walk through it yet

        for dx, dy in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
            nx, ny = x + dx, y + dy
            if (nx, ny) not in grid or grid[(nx, ny)] == "#":
                continue

            cell = grid[(nx, ny)]
            # If it's a door and we don't have the key, we can't pass
            if "A" <= cell <= "Z" and cell.lower() not in current_keys:
                continue

            if (nx, ny) not in visited:
                visited.add((nx, ny))
                queue.append(((nx, ny), dist + 1))
    return reachable
