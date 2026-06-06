import { findInvalid } from "./helpers.js";

export function part1(puzzleInput) {
  return findInvalid(puzzleInput.map(Number), 25);
}
