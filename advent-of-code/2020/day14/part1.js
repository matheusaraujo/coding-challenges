export function part1(puzzleInput) {
  const mem = {};
  let orMask = 0n, andMask = 0n;
  for (const line of puzzleInput) {
    if (line.startsWith("mask")) {
      const m = line.split(" = ")[1];
      orMask = BigInt("0b" + m.replace(/X/g, "0"));
      andMask = BigInt("0b" + m.replace(/X/g, "1"));
    } else {
      const [, addr, val] = line.match(/mem\[(\d+)\] = (\d+)/);
      mem[addr] = (BigInt(val) | orMask) & andMask;
    }
  }
  return Number(Object.values(mem).reduce((a, b) => a + b, 0n));
}
