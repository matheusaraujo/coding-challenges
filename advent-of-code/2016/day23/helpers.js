export class Computer {
  constructor(instructions, initialState = {}) {
    this.instructions = instructions.map((instruction) =>
      instruction.split(" ")
    );
    this.registers = { a: 0, b: 0, c: 0, d: 0, ...initialState };

    this.operations = {
      cpy: (arg1, arg2) => {
        if (isNaN(arg2)) {
          this.registers[arg2] = isNaN(arg1)
            ? this.registers[arg1]
            : parseInt(arg1);
        }
      },
      inc: (arg1) => {
        if (isNaN(arg1)) this.registers[arg1]++;
      },
      dec: (arg1) => {
        if (isNaN(arg1)) this.registers[arg1]--;
      },
      jnz: (arg1, arg2) => {
        const value = isNaN(arg1) ? this.registers[arg1] : parseInt(arg1);
        const offset = isNaN(arg2) ? this.registers[arg2] : parseInt(arg2);
        return value !== 0 ? offset : 1;
      },
      tgl: (arg1, memPointer) => {
        const offset = isNaN(arg1) ? this.registers[arg1] : parseInt(arg1);
        const targetIdx = memPointer + offset;

        if (targetIdx >= 0 && targetIdx < this.instructions.length) {
          const targetInstr = this.instructions[targetIdx];

          if (targetInstr.length === 2) {
            targetInstr[0] = targetInstr[0] === "inc" ? "dec" : "inc";
          } else if (targetInstr.length === 3) {
            targetInstr[0] = targetInstr[0] === "jnz" ? "cpy" : "jnz";
          }
        }
      },
    };
  }

  execute() {
    let memPointer = 0;

    while (memPointer < this.instructions.length) {
      if (this.canOptimizeMultiplication(memPointer)) {
        const [inc, dec1, _jnz1, dec2] = this.instructions.slice(
          memPointer,
          memPointer + 5,
        );

        const targetReg = inc[1];
        const countReg1 = dec1[1];
        const countReg2 = dec2[1];

        this.registers[targetReg] += this.registers[countReg1] *
          this.registers[countReg2];
        this.registers[countReg1] = 0;
        this.registers[countReg2] = 0;

        memPointer += 5;
        continue;
      }

      const [instruction, arg1, arg2] = this.instructions[memPointer];
      let jumpOffset;

      if (instruction === "tgl") {
        this.operations.tgl(arg1, memPointer);
      } else if (this.operations[instruction]) {
        jumpOffset = this.operations[instruction](arg1, arg2);
      }

      memPointer += (jumpOffset !== undefined) ? jumpOffset : 1;
    }

    return this.registers.a;
  }

  canOptimizeMultiplication(ptr) {
    if (ptr + 4 >= this.instructions.length) return false;
    const [i0, i1, i2, i3, i4] = this.instructions.slice(ptr, ptr + 5);

    return (
      i0[0] === "inc" &&
      i1[0] === "dec" &&
      i2[0] === "jnz" && i2[2] === "-2" && i2[1] === i1[1] &&
      i3[0] === "dec" &&
      i4[0] === "jnz" && i4[2] === "-5" && i4[1] === i3[1]
    );
  }
}
