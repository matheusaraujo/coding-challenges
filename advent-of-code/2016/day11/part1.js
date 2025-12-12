import { parseInput, solve } from "./helpers.js";

export function part1(puzzleInput) {
  const { floors } = parseInput(puzzleInput);
  const state = [0, floors];
  return solve(state);
}
