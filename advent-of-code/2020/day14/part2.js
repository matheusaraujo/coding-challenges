export function part2(puzzleInput) {
  const mem = {};
  let mask = "";
  for (const line of puzzleInput) {
    if (line.startsWith("mask")) {
      mask = line.split(" = ")[1];
    } else {
      const [, addr, val] = line.match(/mem\[(\d+)\] = (\d+)/);
      const base = BigInt(addr) | BigInt("0b" + mask.replace(/X/g, "0"));
      const floats = [...mask.matchAll(/X/g)].map((m) => 35 - m.index);
      const count = 1 << floats.length;
      for (let i = 0; i < count; i++) {
        let a = base;
        for (let b = 0; b < floats.length; b++) {
          const bit = BigInt(floats[b]);
          a = i & (1 << b) ? a | (1n << bit) : a & ~(1n << bit);
        }
        mem[a] = BigInt(val);
      }
    }
  }
  return Number(Object.values(mem).reduce((a, b) => a + b, 0n));
}
