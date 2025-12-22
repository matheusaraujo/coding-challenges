export class Computer {
  constructor(instructions, initialState = {}) {
    this.instructions = instructions.map((line) => line.split(" "));
    this.registers = { a: 0, b: 0, c: 0, d: 0, ...initialState };

    this.output = [];

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
        const val = isNaN(arg1) ? this.registers[arg1] : parseInt(arg1);
        const offset = isNaN(arg2) ? this.registers[arg2] : parseInt(arg2);
        return val !== 0 ? offset : 1;
      },
      tgl: (arg1, ptr) => {
        const offset = isNaN(arg1) ? this.registers[arg1] : parseInt(arg1);
        const targetIdx = ptr + offset;
        if (targetIdx >= 0 && targetIdx < this.instructions.length) {
          const instr = this.instructions[targetIdx];
          if (instr.length === 2) {
            instr[0] = instr[0] === "inc" ? "dec" : "inc";
          } else {
            instr[0] = instr[0] === "jnz" ? "cpy" : "jnz";
          }
        }
      },
      out: (arg1) => {
        const val = isNaN(arg1) ? this.registers[arg1] : parseInt(arg1);
        this.output.push(val);
      },
    };
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

  execute(maxOutputLength = Infinity) {
    let ptr = 0;

    while (ptr < this.instructions.length) {
      if (this.canOptimizeMultiplication(ptr)) {
        const [inc, dec1, , dec2] = this.instructions.slice(ptr, ptr + 5);
        this.registers[inc[1]] += this.registers[dec1[1]] *
          this.registers[dec2[1]];
        this.registers[dec1[1]] = 0;
        this.registers[dec2[1]] = 0;
        ptr += 5;
        continue;
      }

      const [op, arg1, arg2] = this.instructions[ptr];

      if (op === "out") {
        this.operations.out(arg1);

        const lastIdx = this.output.length - 1;
        if (this.output[lastIdx] !== (lastIdx % 2)) return false;

        if (this.output.length >= maxOutputLength) return true;
        ptr++;
      } else if (op === "tgl") {
        this.operations.tgl(arg1, ptr);
        ptr++;
      } else if (this.operations[op]) {
        const result = this.operations[op](arg1, arg2);
        ptr += (op === "jnz" && result !== undefined) ? result : 1;
      } else {
        ptr++;
      }
    }
    return this.registers.a;
  }
}
