from collections import deque

from part1 import get_map


def part2(puzzle_input):
    maze, start_pos, _ = get_map(puzzle_input)

    queue = deque([(start_pos, 0)])
    visited = {start_pos}
    max_time = 0

    moves = [(0, -1), (0, 1), (-1, 0), (1, 0)]

    while queue:
        (x, y), time = queue.popleft()
        max_time = max(max_time, time)

        for dx, dy in moves:
            nx, ny = x + dx, y + dy
            # If it's a valid path (1) or the oxygen tank itself (2) and not visited
            if maze.get((nx, ny), 0) != 0 and (nx, ny) not in visited:
                visited.add((nx, ny))
                queue.append(((nx, ny), time + 1))

    return max_time
