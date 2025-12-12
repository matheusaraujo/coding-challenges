import { findKeyIndex } from "./helpers.js";

export function part2(puzzleInput) {
  return findKeyIndex(puzzleInput[0], 64, 2017);
}
