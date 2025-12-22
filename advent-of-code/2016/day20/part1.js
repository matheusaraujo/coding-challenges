import { parseInput } from "./helpers.js";

export function part1(puzzleInput) {
  const ranges = parseInput(puzzleInput);
  let minIP = 0;
  for (const range of ranges) {
    if (minIP < range.start) break;
    minIP = Math.max(minIP, range.end + 1);
  }
  return minIP;
}
