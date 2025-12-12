export function part1(puzzleInput) {
  const map = { A: 0, B: 1, C: 3 };
  return puzzleInput[0]
    .split("")
    .map((x) => map[x])
    .reduce((x, y) => x + y);
}
