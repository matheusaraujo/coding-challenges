import { parse, run } from "./helpers.js";

export function part1(puzzleInput) {
  return run(parse(puzzleInput)).acc;
}
