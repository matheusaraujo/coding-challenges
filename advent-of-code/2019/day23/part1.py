from helpers import Intcode


def part1(puzzle_input):
    if isinstance(puzzle_input, str):
        program = [int(x) for x in puzzle_input.strip().split(",")]
    else:
        program = [int(x) for x in puzzle_input[0].strip().split(",")]

    computers = [Intcode(program, i) for i in range(50)]

    while True:
        for i in range(50):
            status, packet = computers[i].run_step()

            if status == "WAIT":
                # Provide -1 if no packets are in queue
                computers[i].queue.append(-1)
            elif status == "SEND":
                dest, x, y = packet
                if dest == 255:
                    return y
                computers[dest].queue.append(x)
                computers[dest].queue.append(y)
