// Chinese Remainder Theorem via successive sieving
export function part2(puzzleInput) {
  const buses = puzzleInput[1].split(",").map((x, i) =>
    x === "x" ? null : { id: BigInt(x), offset: BigInt(i) }
  ).filter(Boolean);

  let t = 0n, step = 1n;
  for (const { id, offset } of buses) {
    while ((t + offset) % id !== 0n) t += step;
    step *= id;
  }
  return Number(t);
}
