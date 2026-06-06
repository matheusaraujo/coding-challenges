export function parse(puzzleInput) {
  return puzzleInput.map((line) => {
    const [op, arg] = line.split(" ");
    return { op, arg: Number(arg) };
  });
}

export function run(program) {
  let acc = 0, ip = 0;
  const visited = new Set();
  while (ip < program.length) {
    if (visited.has(ip)) return { acc, loop: true };
    visited.add(ip);
    const { op, arg } = program[ip];
    if (op === "acc") {
      acc += arg;
      ip++;
    } else if (op === "jmp") ip += arg;
    else ip++;
  }
  return { acc, loop: false };
}
