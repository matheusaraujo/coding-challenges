import { findKeyIndex } from "./helpers.js";

export function part1(puzzleInput) {
  return findKeyIndex(puzzleInput[0], 64);
}
