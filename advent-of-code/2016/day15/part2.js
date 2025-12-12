import { findFirstValidTime, parseInput } from "./helpers.js";

export function part2(puzzleInput) {
  const discs = parseInput(puzzleInput);
  discs.push([11, 0]);
  return findFirstValidTime(discs);
}
