import { findCorners, parseTiles } from "./helpers.js";

export function part1(puzzleInput) {
  return findCorners(parseTiles(puzzleInput)).reduce((a, b) => a * b, 1);
}
