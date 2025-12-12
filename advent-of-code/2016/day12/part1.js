import { Computer } from "./helpers.js";

export function part1(puzzleInput) {
  return new Computer(puzzleInput).execute();
}
