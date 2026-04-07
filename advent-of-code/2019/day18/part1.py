import collections
import heapq


# pylint: disable=too-many-locals,too-many-branches,too-many-statements,invalid-name
def part1(puzzle_input):
    # Parse the grid
    if isinstance(puzzle_input, str):
        lines = puzzle_input.strip().split("\n")
    else:
        lines = [line.strip() for line in puzzle_input if line.strip()]

    grid = {}
    keys = {}
    start_pos = None
    for y, line in enumerate(lines):
        for x, char in enumerate(line):
            grid[x, y] = char
            if char == "@":
                start_pos = (x, y)
            elif "a" <= char <= "z":
                keys[char] = (x, y)

    # 1. Pre-calculate distances between all points of interest (@ and all keys)
    # This turns the map into a weighted graph.
    points = {"@": start_pos, **keys}
    graph = collections.defaultdict(dict)

    for name, pos in points.items():
        queue = collections.deque([(pos, 0, 0)])  # pos, dist, doors_mask
        visited = {pos}
        while queue:
            (c_x, c_y), dist, doors = queue.popleft()

            char = grid[c_x, c_y]
            if char != name and "a" <= char <= "z":
                graph[name][char] = (dist, doors)
                # We continue BFS because we might find other keys behind this one

            for dx, dy in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
                nx, ny = c_x + dx, c_y + dy
                if (nx, ny) not in grid or grid[nx, ny] == "#":
                    continue
                if (nx, ny) in visited:
                    continue

                visited.add((nx, ny))
                new_doors = doors
                if "A" <= grid[nx, ny] <= "Z":
                    # Store doors as a bitmask (A=1, B=2, C=4, etc.)
                    new_doors |= 1 << (ord(grid[nx, ny]) - ord("A"))

                queue.append(((nx, ny), dist + 1, new_doors))

    # 2. Dijkstra / BFS over the Graph
    target_mask = (1 << len(keys)) - 1
    # State: (current_key, collected_mask)
    # Using a list for the priority queue: (distance, current_key, mask)
    p_q = [(0, "@", 0)]
    visited_states = {}  # (current_key, mask) -> min_dist

    while p_q:
        dist, current, mask = heapq.heappop(p_q)

        if mask == target_mask:
            return dist

        state = (current, mask)
        if state in visited_states and visited_states[state] <= dist:
            continue
        visited_states[state] = dist

        for next_key, (d_dist, d_doors) in graph[current].items():
            key_bit = 1 << (ord(next_key) - ord("a"))

            # Can we reach this key? (Do we have all required doors open?)
            if (mask & d_doors) == d_doors:
                new_mask = mask | key_bit
                if (next_key, new_mask) not in visited_states or visited_states[
                    next_key, new_mask
                ] > dist + d_dist:
                    heapq.heappush(p_q, (dist + d_dist, next_key, new_mask))

    return 0
