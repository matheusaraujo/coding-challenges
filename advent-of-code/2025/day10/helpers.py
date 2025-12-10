import re
from typing import List


# pylint: disable=too-few-public-methods
class Machine:
    def __init__(
        self,
        target_state: str = "",
        buttons: List[List[int]] = None,
        joltage: List[int] = None,
    ):
        self.target_state = target_state
        self.buttons = buttons if buttons is not None else []
        self.joltage = joltage if joltage is not None else []


def parse_input(puzzle_input: List[str]) -> List[Machine]:
    machines: List[Machine] = []
    pattern = re.compile(r"\[([^\]]+)\]|\(([^\)]+)\)|\{([^\}]+)\}")

    for line in puzzle_input:
        m = Machine()
        matches = pattern.findall(line)
        if matches and matches[0][0]:
            m.target_state = matches[0][0].strip()
            matches = matches[1:]
        else:
            continue
        for _group1, group2, group3 in matches:
            if group2:
                raw_nums = [n.strip() for n in group2.split(",")]
                button_set = [int(n) for n in raw_nums if n]
                if button_set:
                    m.buttons.append(button_set)

            elif group3:
                raw_nums = [n.strip() for n in group3.split(",")]
                m.joltage = [int(n) for n in raw_nums if n]

        machines.append(m)

    return machines
