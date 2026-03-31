from helpers import Intcode


def part1(puzzle_input):
    # Handle both string and list inputs safely
    if isinstance(puzzle_input, list):
        raw_data = "".join(puzzle_input)
    else:
        raw_data = puzzle_input

    program = [int(x) for x in raw_data.strip().split(",")]

    robot = Intcode(program)
    hull = {}
    pos = (0, 0)
    direction = 0  # 0: Up, 1: Right, 2: Down, 3: Left
    moves = [(0, -1), (1, 0), (0, 1), (-1, 0)]

    while not robot.halted:
        current_color = hull.get(pos, 0)
        paint_color = robot.run([current_color])

        # If the robot halts before providing the second output
        if paint_color is None:
            break

        turn = robot.run([])
        if turn is None:
            break

        hull[pos] = paint_color
        direction = (direction + (1 if turn == 1 else -1)) % 4
        pos = (pos[0] + moves[direction][0], pos[1] + moves[direction][1])

    return len(hull)
