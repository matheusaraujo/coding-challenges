import { parseBlack } from "./helpers.js";

export function part1(puzzleInput) {
  return parseBlack(puzzleInput).size;
}
