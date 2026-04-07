import collections
import heapq


# pylint: disable=too-many-locals,too-many-branches,too-many-statements,invalid-name
def part2(puzzle_input):
    # 1. Parse and modify grid for Part 2
    if isinstance(puzzle_input, str):
        lines = [
            list(line.strip())
            for line in puzzle_input.strip().split("\n")
            if line.strip()
        ]
    else:
        lines = [list(line.strip()) for line in puzzle_input if line.strip()]

    R, C = len(lines), len(lines[0])
    grid = {}
    keys = {}
    starts = []

    # Find the original '@'
    cx, cy = 0, 0
    for y in range(R):
        for x in range(C):
            if lines[y][x] == "@":
                cx, cy = x, y
            grid[x, y] = lines[y][x]

    # Modify the center for 4 robots
    for dx in [-1, 0, 1]:
        for dy in [-1, 0, 1]:
            grid[cx + dx, cy + dy] = "#"

    # New starting positions (robots 0, 1, 2, 3)
    start_positions = [
        (cx - 1, cy - 1),
        (cx + 1, cy - 1),
        (cx - 1, cy + 1),
        (cx + 1, cy + 1),
    ]
    for i, pos in enumerate(start_positions):
        grid[pos] = str(i)  # Name robots '0', '1', '2', '3'
        starts.append(str(i))

    # Catalog all keys
    for (x, y), char in grid.items():
        if "a" <= char <= "z":
            keys[char] = (x, y)

    # 2. Pre-calculate distances between all points (@0-3 and all keys)
    points = {**{str(i): pos for i, pos in enumerate(start_positions)}, **keys}
    graph = collections.defaultdict(dict)

    for name, pos in points.items():
        queue = collections.deque([(pos, 0, 0)])
        visited = {pos}
        while queue:
            (cx, cy), dist, doors = queue.popleft()
            char = grid[cx, cy]
            if char != name and "a" <= char <= "z":
                graph[name][char] = (dist, doors)

            for dx, dy in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
                nx, ny = cx + dx, cy + dy
                if (nx, ny) not in grid or grid[nx, ny] == "#":
                    continue
                if (nx, ny) in visited:
                    continue
                visited.add((nx, ny))
                new_doors = doors
                if "A" <= grid[nx, ny] <= "Z":
                    new_doors |= 1 << (ord(grid[nx, ny]) - ord("A"))
                queue.append(((nx, ny), dist + 1, new_doors))

    # 3. Dijkstra with 4 robot positions
    target_mask = (1 << len(keys)) - 1
    # State: (tuple_of_4_current_locations, collected_mask)
    initial_robots = ("0", "1", "2", "3")
    pq = [(0, initial_robots, 0)]
    visited_states = {}

    while pq:
        dist, robots, mask = heapq.heappop(pq)

        if mask == target_mask:
            return dist

        state = (robots, mask)
        if state in visited_states and visited_states[state] <= dist:
            continue
        visited_states[state] = dist

        # Try moving each robot
        for i in range(4):
            current_loc = robots[i]
            # Which keys can THIS robot reach?
            for next_key, (d_dist, d_doors) in graph[current_loc].items():
                key_bit = 1 << (ord(next_key) - ord("a"))

                # If we don't have this key yet and we have the doors for it
                if not (mask & key_bit) and (mask & d_doors) == d_doors:
                    new_mask = mask | key_bit
                    new_robots = list(robots)
                    new_robots[i] = next_key
                    new_robots = tuple(new_robots)

                    if (new_robots, new_mask) not in visited_states or visited_states[
                        new_robots, new_mask
                    ] > dist + d_dist:
                        heapq.heappush(pq, (dist + d_dist, new_robots, new_mask))
    return 0
