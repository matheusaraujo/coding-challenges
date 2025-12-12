import { parseInput } from "./helpers.js";

export function part1(puzzleInput) {
  const { left, right } = parseInput(puzzleInput);
  return left.reduce((sum, l, index) => sum + Math.abs(l - right[index]), 0);
}
