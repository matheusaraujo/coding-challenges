# part2.py

from helpers import IntcodeComputer, get_grid, parse_program

DIRS = {
    "^": (0, -1),
    "v": (0, 1),
    "<": (-1, 0),
    ">": (1, 0),
}

LEFT = {
    "^": "<",
    "<": "v",
    "v": ">",
    ">": "^",
}

RIGHT = {
    "^": ">",
    ">": "v",
    "v": "<",
    "<": "^",
}


def find_robot(grid):
    for y, row in enumerate(grid):
        for x, char in enumerate(row):
            if char in DIRS:
                return x, y, char
    raise ValueError("Robot not found")  # fixes inconsistent return


def is_scaffold(grid, x, y):
    return 0 <= y < len(grid) and 0 <= x < len(grid[0]) and grid[y][x] == "#"


def build_path(grid):
    x, y, direction = find_robot(grid)
    path = []

    while True:
        dx, dy = DIRS[direction]
        steps = 0

        while is_scaffold(grid, x + dx, y + dy):
            x += dx
            y += dy
            steps += 1

        if steps > 0:
            path.append(str(steps))

        left_dir = LEFT[direction]
        left_dx, left_dy = DIRS[left_dir]

        right_dir = RIGHT[direction]
        right_dx, right_dy = DIRS[right_dir]

        if is_scaffold(grid, x + left_dx, y + left_dy):
            path.append("L")
            direction = left_dir
        elif is_scaffold(grid, x + right_dx, y + right_dy):
            path.append("R")
            direction = right_dir
        else:
            break

    return path


def to_string(seq):
    return ",".join(seq)


def fits(seq):
    return len(to_string(seq)) <= 20


def compress(path):
    for a_len in range(2, 11):
        func_a = path[:a_len]
        if not fits(func_a):
            continue

        def replace(seq, pattern, label):
            s = ",".join(seq)
            p = ",".join(pattern)
            return s.replace(p, label)

        seq_after_a = replace(path, func_a, "A").split(",")

        for b_start, token in enumerate(seq_after_a):
            if token in "ABC":
                continue

            for b_len in range(2, 11):
                func_b = seq_after_a[b_start : b_start + b_len]
                if not fits(func_b):
                    continue

                seq_after_b = replace(seq_after_a, func_b, "B").split(",")

                for c_start, token_c in enumerate(seq_after_b):
                    if token_c in "ABC":
                        continue

                    for c_len in range(2, 11):
                        func_c = seq_after_b[c_start : c_start + c_len]
                        if not fits(func_c):
                            continue

                        seq_after_c = replace(seq_after_b, func_c, "C").split(",")

                        if all(x in ["A", "B", "C"] for x in seq_after_c):
                            return seq_after_c, func_a, func_b, func_c

    raise ValueError("No compression found")


def part2(puzzle_input):
    program = parse_program(puzzle_input)

    computer = IntcodeComputer(program[:])
    output = computer.run()
    grid = get_grid(output)

    path = build_path(grid)

    main_routine, func_a, func_b, func_c = compress(path)

    program[0] = 2
    computer = IntcodeComputer(program)

    def encode(line):
        return [ord(c) for c in line] + [10]

    inputs = []
    inputs += encode(to_string(main_routine))
    inputs += encode(to_string(func_a))
    inputs += encode(to_string(func_b))
    inputs += encode(to_string(func_c))
    inputs += encode("n")

    output = computer.run(inputs)

    return output[-1]
