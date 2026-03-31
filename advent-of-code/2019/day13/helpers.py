from collections import defaultdict


# pylint: disable=too-few-public-methods
class Intcode:
    def __init__(self, program):
        self.memory = defaultdict(int, dict(enumerate(program)))
        self.pointer = 0
        self.relative_base = 0
        self.halted = False

    def __get_param(self, instruction, idx):
        mode = int(instruction[3 - idx])
        val = self.memory[self.pointer + idx]
        if mode == 0:
            return self.memory[val]  # Position
        if mode == 1:
            return val  # Immediate
        if mode == 2:
            return self.memory[self.relative_base + val]  # Relative
        return 0

    def __get_addr(self, instruction, idx):
        mode = int(instruction[3 - idx])
        val = self.memory[self.pointer + idx]
        return self.relative_base + val if mode == 2 else val

    # pylint: disable=too-many-branches
    def run(self, inputs):
        while self.pointer in self.memory:
            instruction = str(self.memory[self.pointer]).zfill(5)
            opcode = int(instruction[-2:])

            if opcode == 99:
                self.halted = True
            elif opcode == 1:  # Add
                self.memory[self.__get_addr(instruction, 3)] = self.__get_param(
                    instruction, 1
                ) + self.__get_param(instruction, 2)
                self.pointer += 4
            elif opcode == 2:  # Mul
                self.memory[self.__get_addr(instruction, 3)] = self.__get_param(
                    instruction, 1
                ) * self.__get_param(instruction, 2)
                self.pointer += 4
            elif opcode == 3:  # Input
                if not inputs:
                    return None  # Wait for input
                self.memory[self.__get_addr(instruction, 1)] = inputs.pop(0)
                self.pointer += 2
            elif opcode == 4:  # Output
                res = self.__get_param(instruction, 1)
                self.pointer += 2
                return res
            elif opcode == 5:  # Jump-if-true
                self.pointer = (
                    self.__get_param(instruction, 2)
                    if self.__get_param(instruction, 1) != 0
                    else self.pointer + 3
                )
            elif opcode == 6:  # Jump-if-false
                self.pointer = (
                    self.__get_param(instruction, 2)
                    if self.__get_param(instruction, 1) == 0
                    else self.pointer + 3
                )
            elif opcode == 7:  # Less than
                self.memory[self.__get_addr(instruction, 3)] = (
                    1
                    if self.__get_param(instruction, 1)
                    < self.__get_param(instruction, 2)
                    else 0
                )
                self.pointer += 4
            elif opcode == 8:  # Equals
                self.memory[self.__get_addr(instruction, 3)] = (
                    1
                    if self.__get_param(instruction, 1)
                    == self.__get_param(instruction, 2)
                    else 0
                )
                self.pointer += 4
            elif opcode == 9:  # Relative base offset
                self.relative_base += self.__get_param(instruction, 1)
                self.pointer += 2

            if self.halted:
                break

        return None
