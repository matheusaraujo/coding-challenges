import { parse, run } from "./helpers.js";

export function part2(puzzleInput) {
  const program = parse(puzzleInput);
  for (let i = 0; i < program.length; i++) {
    const { op } = program[i];
    if (op === "acc") continue;
    const patched = program.map((instr, j) =>
      j === i ? { ...instr, op: op === "jmp" ? "nop" : "jmp" } : instr
    );
    const { acc, loop } = run(patched);
    if (!loop) return acc;
  }
}
