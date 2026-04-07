import collections


class Intcode:
    """A virtual machine for executing Intcode programs."""

    def __init__(self, program):
        self.memory = collections.defaultdict(int, enumerate(program))
        self.pointer = 0
        self.relative_base = 0

    def get_memory_size(self):
        """Returns the number of initialized memory addresses."""
        return len(self.memory)

    # pylint: disable=too-many-branches
    def run(self, inputs):
        """Executes the program with the given inputs."""
        input_ptr = 0
        outputs = []
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
                if input_ptr >= len(inputs):
                    return outputs
                self.memory[get_addr(1)] = inputs[input_ptr]
                input_ptr += 1
                self.pointer += 2
            elif opcode == 4:  # Output
                outputs.append(self.memory[get_addr(1)])
                self.pointer += 2
            elif opcode == 5:  # Jump-if-true
                if self.memory[get_addr(1)] != 0:
                    self.pointer = self.memory[get_addr(2)]
                else:
                    self.pointer += 3
            elif opcode == 6:  # Jump-if-false
                if self.memory[get_addr(1)] == 0:
                    self.pointer = self.memory[get_addr(2)]
                else:
                    self.pointer += 3
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
        return outputs


def run_springscript(puzzle_input, script):
    """Parses input and runs the springdroid script."""
    if isinstance(puzzle_input, str):
        program = [int(x) for x in puzzle_input.strip().split(",")]
    else:
        program = [int(x) for x in puzzle_input[0].strip().split(",")]

    ascii_input = [ord(char) for char in script]
    intcode_vm = Intcode(program)
    output = intcode_vm.run(ascii_input)

    if output and output[-1] > 255:
        return output[-1]

    return None
