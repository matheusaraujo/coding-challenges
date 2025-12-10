from collections import deque

from helpers import Machine, parse_input


def part1(puzzle_input):
    machines = parse_input(puzzle_input)
    result = 0

    for machine in machines:
        result += configure(machine)

    return result


def _next_state(state, buttons) -> str:
    state_list = list(state)
    for pos in buttons:
        if state_list[pos] == ".":
            state_list[pos] = "#"
        elif state_list[pos] == "#":
            state_list[pos] = "."
    return "".join(state_list)


def configure(machine: Machine) -> int:
    target_state = machine.target_state
    buttons = machine.buttons
    state_length = len(target_state)

    start_state = "." * state_length

    queue: deque[tuple[str, int]] = deque([(start_state, 0)])

    visited: set[str] = {start_state}

    while queue:
        current_state, dist = queue.popleft()

        for button_set in buttons:
            next_state = _next_state(current_state, button_set)

            if next_state == target_state:
                return dist + 1

            if next_state not in visited:
                visited.add(next_state)
                queue.append((next_state, dist + 1))

    return -1
