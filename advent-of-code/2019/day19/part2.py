from helpers import check_pos


def part2(puzzle_input):
    if isinstance(puzzle_input, str):
        program = [int(x) for x in puzzle_input.strip().split(",")]
    else:
        program = [int(x) for x in puzzle_input[0].strip().split(",")]

    # We need to find a 100x100 square.
    # To optimize, we track the 'bottom-left' corner of the beam as we go down.
    # If the point (x, y) is the bottom-left of the square, then the
    # top-right of the square is at (x + 99, y - 99).

    x = 0
    y = 100  # Start a bit lower to avoid the narrow origin

    while True:
        # Find the start of the beam for this Y
        while check_pos(program, x, y) == 0:
            x += 1

        # Check if the top-right corner of a 100x100 square starting at (x, y)
        # is also within the beam.
        if check_pos(program, x + 99, y - 99) == 1:
            # Found it! The problem asks for x * 10000 + y of the TOP-LEFT corner.
            # Our current x is the left side, and (y - 99) is the top side.
            return x * 10000 + (y - 99)

        y += 1
