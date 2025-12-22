import { Computer } from "./helpers.js";

export function part2(puzzleInput) {
  return new Computer(puzzleInput, { a: 12 }).execute();
}
