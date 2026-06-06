import { parse, simulate } from "./helpers.js";

export function part2(puzzleInput) {
  return simulate(parse(puzzleInput, 4), 4, 6);
}
