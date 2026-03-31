from helpers import Intcode


def part2(puzzle_input):
    raw_data = "".join(puzzle_input) if isinstance(puzzle_input, list) else puzzle_input
    program = [int(x) for x in raw_data.strip().split(",")]

    # Set to 'play for free'
    program[0] = 2
    computer = Intcode(program)

    score = 0
    paddle_x = 0
    ball_x = 0

    # Game Loop
    while not computer.halted:
        # Provide joystick input based on ball position relative to paddle
        joystick = 0
        if ball_x < paddle_x:
            joystick = -1
        elif ball_x > paddle_x:
            joystick = 1

        # Run until we get an output or halt
        out_x = computer.run([joystick])
        if out_x is None:
            break  # Program halted

        out_y = computer.run([])
        out_val = computer.run([])

        if out_x == -1 and out_y == 0:
            score = out_val
        elif out_val == 3:  # Paddle
            paddle_x = out_x
        elif out_val == 4:  # Ball
            ball_x = out_x

    return score
