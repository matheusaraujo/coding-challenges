import { Computer } from "./helpers.js";

export function part1(puzzleInput) {
  return new Computer(puzzleInput, { a: 7 }).execute();
}
