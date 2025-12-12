import { findFirstValidTime, parseInput } from "./helpers.js";

export function part1(puzzleInput) {
  return findFirstValidTime(parseInput(puzzleInput));
}
