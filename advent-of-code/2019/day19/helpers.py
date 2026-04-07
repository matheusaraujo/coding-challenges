import collections


# pylint: disable=too-few-public-methods
class Intcode:
    def __init__(self, program):
        self.memory = collections.defaultdict(int, enumerate(program))
        self.pointer = 0
        self.relative_base = 0

    def run(self, inputs):
        input_ptr = 0
        while True:
            instruction = str(self.memory[self.pointer]).zfill(5)
            opcode = int(instruction[-2:])
            modes = [int(x) for x in instruction[:3][::-1]]

            def get_addr(i):
                if modes[i - 1] == 0:
                    return self.memory[self.pointer + i]
                if modes[i - 1] == 1:
                    return self.pointer + i
                if modes[i - 1] == 2:
                    return self.relative_base + self.memory[self.pointer + i]
                return None

            if opcode == 99:
                break

            if opcode == 1:  # Add
                self.memory[get_addr(3)] = (
                    self.memory[get_addr(1)] + self.memory[get_addr(2)]
                )
                self.pointer += 4
            elif opcode == 2:  # Mult
                self.memory[get_addr(3)] = (
                    self.memory[get_addr(1)] * self.memory[get_addr(2)]
                )
                self.pointer += 4
            elif opcode == 3:  # Input
                self.memory[get_addr(1)] = inputs[input_ptr]
                input_ptr += 1
                self.pointer += 2
            elif opcode == 4:  # Output
                res = self.memory[get_addr(1)]
                self.pointer += 2
                return res
            elif opcode == 5:  # Jump-if-true
                self.pointer = (
                    self.memory[get_addr(2)]
                    if self.memory[get_addr(1)] != 0
                    else self.pointer + 3
                )
            elif opcode == 6:  # Jump-if-false
                self.pointer = (
                    self.memory[get_addr(2)]
                    if self.memory[get_addr(1)] == 0
                    else self.pointer + 3
                )
            elif opcode == 7:  # Less than
                self.memory[get_addr(3)] = (
                    1 if self.memory[get_addr(1)] < self.memory[get_addr(2)] else 0
                )
                self.pointer += 4
            elif opcode == 8:  # Equals
                self.memory[get_addr(3)] = (
                    1 if self.memory[get_addr(1)] == self.memory[get_addr(2)] else 0
                )
                self.pointer += 4
            elif opcode == 9:  # Relative Base
                self.relative_base += self.memory[get_addr(1)]
                self.pointer += 2

        return None  # Explicit return to satisfy R1710


def check_pos(program, x, y):
    """Returns 1 if the tractor beam is at (x, y), 0 otherwise."""
    return Intcode(program).run([x, y])
