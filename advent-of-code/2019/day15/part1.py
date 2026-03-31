from collections import deque

from helpers import Intcode


# pylint: disable=too-many-locals
def get_map(puzzle_input):
    raw = "".join(puzzle_input) if isinstance(puzzle_input, list) else puzzle_input
    prog = [int(x) for x in raw.strip().split(",")]

    # directions: 1=N, 2=S, 3=W, 4=E
    moves = {1: (0, -1), 2: (0, 1), 3: (-1, 0), 4: (1, 0)}
    queue = deque([(0, 0, Intcode(prog), 0)])
    visited = {(0, 0)}
    maze = {(0, 0): 1}
    oxygen_pos = None
    oxygen_dist = 0

    while queue:
        x, y, droid, dist = queue.popleft()

        for d, (dx, dy) in moves.items():
            nx, ny = x + dx, y + dy
            if (nx, ny) in visited:
                continue

            new_droid = droid.clone()
            status = new_droid.run([d])

            if status == 0:  # Wall
                maze[(nx, ny)] = 0
                visited.add((nx, ny))
            else:
                maze[(nx, ny)] = status
                visited.add((nx, ny))
                if status == 2:
                    oxygen_pos = (nx, ny)
                    oxygen_dist = dist + 1
                queue.append((nx, ny, new_droid, dist + 1))

    return maze, oxygen_pos, oxygen_dist


def part1(puzzle_input):
    _, _, dist = get_map(puzzle_input)
    return dist
