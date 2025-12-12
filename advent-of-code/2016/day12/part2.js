import { Computer } from "./helpers.js";

export function part2(puzzleInput) {
  return new Computer(puzzleInput, { c: 1 }).execute();
}
