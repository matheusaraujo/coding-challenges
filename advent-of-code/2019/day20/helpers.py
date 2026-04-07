# pylint: disable=too-many-locals
def parse_maze(puzzle_input):
    if isinstance(puzzle_input, str):
        lines = puzzle_input.split("\n")
    else:
        lines = [line.replace("\n", "") for line in puzzle_input]

    grid = {}
    for y, line in enumerate(lines):
        for x, char in enumerate(line):
            grid[(x, y)] = char

    # Optimized: Iterating directly over the dict instead of .keys()
    max_x = max(x for x, y in grid)
    max_y = max(y for x, y in grid)

    portals = {}  # (x, y) -> label
    label_coords = {}  # label -> [(x, y), ...]

    for (x, y), char in grid.items():
        if "A" <= char <= "Z":
            # Look for neighbor '.' and the other half of the label
            for dx, dy in [(0, 1), (1, 0), (0, -1), (-1, 0)]:
                nx, ny = x + dx, y + dy
                if grid.get((nx, ny)) == ".":
                    # The label is either (x-dx, y-dy) + (x, y)
                    # or (x, y) + (x-dx, y-dy) depending on order.
                    other_x, other_y = x - dx, y - dy
                    other_char = grid.get((other_x, other_y))

                    if not other_char or not "A" <= other_char <= "Z":
                        continue

                    # Sort by coordinate to keep label string consistent
                    # (Top-to-Bottom, Left-to-Right)
                    chars = sorted(
                        [((x, y), char), ((other_x, other_y), other_char)],
                        key=lambda p: (p[0][1], p[0][0]),
                    )
                    label = chars[0][1] + chars[1][1]

                    portals[(nx, ny)] = label
                    if label not in label_coords:
                        label_coords[label] = []
                    # Avoid adding the same entrance twice
                    if (nx, ny) not in label_coords[label]:
                        label_coords[label].append((nx, ny))

    return grid, portals, label_coords, max_x, max_y


def is_outer(pos, max_x, max_y):
    x, y = pos
    # Maze path '.' usually starts at index 2 due to the labels and a 1-char border
    return x <= 2 or y <= 2 or x >= max_x - 2 or y >= max_y - 2
