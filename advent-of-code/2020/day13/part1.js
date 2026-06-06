export function part1(puzzleInput) {
  const earliest = Number(puzzleInput[0]);
  const buses = puzzleInput[1].split(",").filter((x) => x !== "x").map(Number);
  const [wait, bus] = buses
    .map((b) => [b - (earliest % b === 0 ? b : earliest % b), b])
    .sort((a, b) => a[0] - b[0])[0];
  return wait * bus;
}
