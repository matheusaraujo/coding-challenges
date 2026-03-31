import copy
from collections import defaultdict


class Intcode:
    def __init__(self, program, memory=None, pointer=0, relative_base=0):
        if memory is not None:
            self.memory = memory
        else:
            self.memory = defaultdict(int, dict(enumerate(program)))
        self.pointer = pointer
        self.relative_base = relative_base
        self.halted = False

    def clone(self):
        return Intcode(
            None, copy.deepcopy(self.memory), self.pointer, self.relative_base
        )

    def run(self, inputs):
        while True:
            inst = str(self.memory[self.pointer]).zfill(5)
            opcode = int(inst[-2:])

            def get_p(i):
                m, v = int(inst[3 - i]), self.memory[self.pointer + i]
                if m == 0:
                    return self.memory[v]
                if m == 1:
                    return v
                return self.memory[self.relative_base + v]

            def get_a(i):
                m, v = int(inst[3 - i]), self.memory[self.pointer + i]
                return self.relative_base + v if m == 2 else v

            if opcode == 99:
                self.halted = True
                return None
            if opcode == 1:
                self.memory[get_a(3)] = get_p(1) + get_p(2)
                self.pointer += 4
            elif opcode == 2:
                self.memory[get_a(3)] = get_p(1) * get_p(2)
                self.pointer += 4
            elif opcode == 3:
                if not inputs:
                    return None
                self.memory[get_a(1)] = inputs.pop(0)
                self.pointer += 2
            elif opcode == 4:
                res = get_p(1)
                self.pointer += 2
                return res
            elif opcode == 5:
                self.pointer = get_p(2) if get_p(1) != 0 else self.pointer + 3
            elif opcode == 6:
                self.pointer = get_p(2) if get_p(1) == 0 else self.pointer + 3
            elif opcode == 7:
                self.memory[get_a(3)] = 1 if get_p(1) < get_p(2) else 0
                self.pointer += 4
            elif opcode == 8:
                self.memory[get_a(3)] = 1 if get_p(1) == get_p(2) else 0
                self.pointer += 4
            elif opcode == 9:
                self.relative_base += get_p(1)
                self.pointer += 2
