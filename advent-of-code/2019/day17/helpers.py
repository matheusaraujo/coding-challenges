# helpers.py

from collections import defaultdict


# pylint: disable=too-few-public-methods
class IntcodeComputer:
    """Minimal Intcode computer for AoC 2019."""

    def __init__(self, program):
        self.mem = defaultdict(int, enumerate(program))
        self.i = 0
        self.rel = 0
        self.output = []
        self.halted = False

    def _get(self, k, modes):
        mode = modes[k - 1]
        if mode == 0:
            return self.mem[self.mem[self.i + k]]
        if mode == 1:
            return self.mem[self.i + k]
        return self.mem[self.mem[self.i + k] + self.rel]

    def _set(self, k, value, modes):
        mode = modes[k - 1]
        if mode == 2:
            self.mem[self.mem[self.i + k] + self.rel] = value
        else:
            self.mem[self.mem[self.i + k]] = value

    # pylint: disable=too-many-branches
    def run(self, inputs=None):
        if inputs is None:
            inputs = []

        inputs = inputs[::-1]

        while True:
            opcode = self.mem[self.i] % 100
            modes = [(self.mem[self.i] // 10 ** k) % 10 for k in range(2, 5)]

            if opcode == 1:
                self._set(3, self._get(1, modes) + self._get(2, modes), modes)
                self.i += 4

            elif opcode == 2:
                self._set(3, self._get(1, modes) * self._get(2, modes), modes)
                self.i += 4

            elif opcode == 3:
                if not inputs:
                    return self.output
                self._set(1, inputs.pop(), modes)
                self.i += 2

            elif opcode == 4:
                self.output.append(self._get(1, modes))
                self.i += 2

            elif opcode == 5:
                if self._get(1, modes) != 0:
                    self.i = self._get(2, modes)
                else:
                    self.i += 3

            elif opcode == 6:
                if self._get(1, modes) == 0:
                    self.i = self._get(2, modes)
                else:
                    self.i += 3

            elif opcode == 7:
                self._set(
                    3,
                    int(self._get(1, modes) < self._get(2, modes)),
                    modes,
                )
                self.i += 4

            elif opcode == 8:
                self._set(
                    3,
                    int(self._get(1, modes) == self._get(2, modes)),
                    modes,
                )
                self.i += 4

            elif opcode == 9:
                self.rel += self._get(1, modes)
                self.i += 2

            elif opcode == 99:
                self.halted = True
                return self.output

            else:
                raise ValueError(f"Unknown opcode: {opcode}")


def parse_program(puzzle_input):
    if isinstance(puzzle_input, list):
        puzzle_input = puzzle_input[0]
    return list(map(int, puzzle_input.strip().split(",")))


def get_grid(output):
    text = "".join(chr(c) for c in output)
    return [list(line) for line in text.strip().split("\n")]
