import collections


# pylint: disable=too-few-public-methods
class Intcode:
    def __init__(self, program, address):
        self.memory = collections.defaultdict(int, enumerate(program))
        self.pointer = 0
        self.relative_base = 0
        self.queue = collections.deque([address])

    # pylint: disable=too-many-branches
    def run_step(self):
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
                return "HALT", outputs

            if opcode == 3:  # Input
                if not self.queue:
                    return "WAIT", outputs
                self.memory[get_addr(1, mode1)] = self.queue.popleft()
                self.pointer += 2
            elif opcode == 4:  # Output
                val = self.memory[get_addr(1, mode1)]
                outputs.append(val)
                self.pointer += 2
                if len(outputs) == 3:
                    return "SEND", outputs
            elif opcode == 1:  # Add
                self.memory[get_addr(3, mode3)] = (
                    self.memory[get_addr(1, mode1)] + self.memory[get_addr(2, mode2)]
                )
                self.pointer += 4
            elif opcode == 2:  # Mult
                self.memory[get_addr(3, mode3)] = (
                    self.memory[get_addr(1, mode1)] * self.memory[get_addr(2, mode2)]
                )
                self.pointer += 4
            elif opcode == 5:  # JNZ
                self.pointer = (
                    self.memory[get_addr(2, mode2)]
                    if self.memory[get_addr(1, mode1)] != 0
                    else self.pointer + 3
                )
            elif opcode == 6:  # JZ
                self.pointer = (
                    self.memory[get_addr(2, mode2)]
                    if self.memory[get_addr(1, mode1)] == 0
                    else self.pointer + 3
                )
            elif opcode == 7:  # LT
                self.memory[get_addr(3, mode3)] = (
                    1
                    if self.memory[get_addr(1, mode1)] < self.memory[get_addr(2, mode2)]
                    else 0
                )
                self.pointer += 4
            elif opcode == 8:  # EQ
                self.memory[get_addr(3, mode3)] = (
                    1
                    if self.memory[get_addr(1, mode1)]
                    == self.memory[get_addr(2, mode2)]
                    else 0
                )
                self.pointer += 4
            elif opcode == 9:  # ARB
                self.relative_base += self.memory[get_addr(1, mode1)]
                self.pointer += 2
