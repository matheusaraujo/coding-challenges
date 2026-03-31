from helpers import Intcode


def parse_ocr(grid_str):
    alphabet = {
        ".##.\n#..#\n#..#\n####\n#..#\n#..#": "A",
        "###.\n#..#\n###.\n#..#\n#..#\n###.": "B",
        ".##.\n#..#\n#...\n#...\n#..#\n.##.": "C",
        "####\n#...\n###.\n#...\n#...\n####": "E",
        "####\n#...\n###.\n#...\n#...\n#...": "F",
        ".##.\n#..#\n#...\n#.##\n#..#\n.###": "G",
        "#..#\n#..#\n####\n#..#\n#..#\n#..#": "H",
        ".###\n..#.\n..#.\n..#.\n..#.\n.###": "I",
        "..##\n...#\n...#\n...#\n#..#\n.##.": "J",
        "#..#\n#.#.\n##..\n#.#.\n#.#.\n#..#": "K",
        "#...\n#...\n#...\n#...\n#...\n####": "L",
        ".##.\n#..#\n#..#\n#..#\n#..#\n.##.": "O",
        "###.\n#..#\n#..#\n###.\n#...\n#...": "P",
        "###.\n#..#\n#..#\n###.\n#.#.\n#..#": "R",
        ".###\n#...\n#...\n.##.\n...#\n###.": "S",
        "####\n..#.\n..#.\n..#.\n..#.\n..#.": "T",
        "#..#\n#..#\n#..#\n#..#\n#..#\n.##.": "U",
        "####\n...#\n..#.\n.#..\n#...\n####": "Z",
    }

    lines = [l for l in grid_str.split("\n") if l.strip()]
    if not lines:
        return "EMPTY_GRID"

    try:
        first_pixel = min(line.find("#") for line in lines if "#" in line)
    except ValueError:
        return "NO_WHITE_PANELS"

    width = len(lines[0])
    result = ""
    for x in range(first_pixel, width, 5):
        char_segment = []
        for y in range(6):
            slice_4 = lines[y][x : x + 4].replace("#", "#").replace(" ", ".")
            char_segment.append(slice_4.ljust(4, "."))

        fingerprint = "\n".join(char_segment)
        if fingerprint.replace(".", "").replace("\n", "") == "":
            continue
        result += alphabet.get(fingerprint, "?")

    return result


def part2(puzzle_input):
    # Combine parsing to reduce local variables (Fixes R0914)
    raw_data = "".join(puzzle_input) if isinstance(puzzle_input, list) else puzzle_input
    program = [int(x) for x in raw_data.strip().split(",")]

    robot = Intcode(program)
    hull, pos, direction = {(0, 0): 1}, (0, 0), 0
    moves = [(0, -1), (1, 0), (0, 1), (-1, 0)]

    while not robot.halted:
        paint = robot.run([hull.get(pos, 0)])
        if paint is None:
            break

        turn = robot.run([])
        if turn is None:
            break

        hull[pos] = paint
        direction = (direction + (1 if turn == 1 else -1)) % 4
        pos = (pos[0] + moves[direction][0], pos[1] + moves[direction][1])

    x_s = [p[0] for p in hull]
    y_s = [p[1] for p in hull]

    grid_rows = []
    for y in range(min(y_s), max(y_s) + 1):
        line = "".join(
            "#" if hull.get((x, y), 0) == 1 else " "
            for x in range(min(x_s), max(x_s) + 1)
        )
        grid_rows.append(line)

    return parse_ocr("\n".join(grid_rows))
