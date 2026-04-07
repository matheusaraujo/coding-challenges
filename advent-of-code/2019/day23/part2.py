from helpers import Intcode


def part2(puzzle_input):
    if isinstance(puzzle_input, str):
        program = [int(x) for x in puzzle_input.strip().split(",")]
    else:
        program = [int(x) for x in puzzle_input[0].strip().split(",")]

    computers = [Intcode(program, i) for i in range(50)]
    nat_x, nat_y = None, None
    last_y_sent_to_zero = None

    while True:
        idle_this_round = True

        for i in range(50):
            # If queue is empty, protocol says provide -1
            if not computers[i].queue:
                computers[i].queue.append(-1)

            status, packet = computers[i].run_step()

            if status == "SEND":
                idle_this_round = False
                dest, x, y = packet
                if dest == 255:
                    nat_x, nat_y = x, y
                else:
                    computers[dest].queue.append(x)
                    computers[dest].queue.append(y)

            # If it's waiting and we just gave it a -1, it's potentially idle.
            # But if it had a packet to send, it's not idle.

        # Check for Global Idle: All queues were empty (except for our -1s)
        # and no one sent anything this round.
        if idle_this_round and all(len(c.queue) == 0 for c in computers):
            if nat_y is not None:
                if nat_y == last_y_sent_to_zero:
                    return nat_y

                last_y_sent_to_zero = nat_y
                computers[0].queue.append(nat_x)
                computers[0].queue.append(nat_y)
