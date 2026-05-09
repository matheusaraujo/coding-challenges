import collections


# pylint: disable=too-few-public-methods
class Intcode:
    """Step-pausing Intcode VM. Run yields outputs and pauses when input is needed."""

    def __init__(self, program):
        self.memory = collections.defaultdict(int, enumerate(program))
        self.pointer = 0
        self.relative_base = 0
        self.inputs = collections.deque()
        self.halted = False

    def feed(self, text):
        for char in text:
            self.inputs.append(ord(char))

    # pylint: disable=too-many-branches
    def run(self):
        outputs = []
        while True:
            instruction = self.memory[self.pointer]
            opcode = instruction % 100
            mode1 = (instruction // 100) % 10
            mode2 = (instruction // 1000) % 10
            mode3 = (instruction // 10000) % 10

            def get_addr(i, mode):
                if mode == 0:
                    return self.memory[self.pointer + i]
                if mode == 1:
                    return self.pointer + i
                if mode == 2:
                    return self.relative_base + self.memory[self.pointer + i]
                return None

            if opcode == 99:
                self.halted = True
                return outputs
            if opcode == 1:
                self.memory[get_addr(3, mode3)] = (
                    self.memory[get_addr(1, mode1)] + self.memory[get_addr(2, mode2)]
                )
                self.pointer += 4
            elif opcode == 2:
                self.memory[get_addr(3, mode3)] = (
                    self.memory[get_addr(1, mode1)] * self.memory[get_addr(2, mode2)]
                )
                self.pointer += 4
            elif opcode == 3:
                if not self.inputs:
                    return outputs
                self.memory[get_addr(1, mode1)] = self.inputs.popleft()
                self.pointer += 2
            elif opcode == 4:
                outputs.append(self.memory[get_addr(1, mode1)])
                self.pointer += 2
            elif opcode == 5:
                self.pointer = (
                    self.memory[get_addr(2, mode2)]
                    if self.memory[get_addr(1, mode1)] != 0
                    else self.pointer + 3
                )
            elif opcode == 6:
                self.pointer = (
                    self.memory[get_addr(2, mode2)]
                    if self.memory[get_addr(1, mode1)] == 0
                    else self.pointer + 3
                )
            elif opcode == 7:
                self.memory[get_addr(3, mode3)] = (
                    1
                    if self.memory[get_addr(1, mode1)] < self.memory[get_addr(2, mode2)]
                    else 0
                )
                self.pointer += 4
            elif opcode == 8:
                self.memory[get_addr(3, mode3)] = (
                    1
                    if self.memory[get_addr(1, mode1)]
                    == self.memory[get_addr(2, mode2)]
                    else 0
                )
                self.pointer += 4
            elif opcode == 9:
                self.relative_base += self.memory[get_addr(1, mode1)]
                self.pointer += 2
