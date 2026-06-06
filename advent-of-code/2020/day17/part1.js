import { parse, simulate } from "./helpers.js";

export function part1(puzzleInput) {
  return simulate(parse(puzzleInput, 3), 3, 6);
}
