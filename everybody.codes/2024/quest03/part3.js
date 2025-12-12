import { solve } from "./helpers.js";

export function part3(puzzleInput) {
  return solve(puzzleInput, [
    [1, 0],
    [-1, 0],
    [0, 1],
    [0, -1],
    [1, 1],
    [1, -1],
    [-1, 1],
    [-1, -1],
  ]);
}
