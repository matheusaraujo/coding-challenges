import { solve } from "./helpers.js";

export function part3(puzzleInput) {
  return solve(puzzleInput, (arr) => Math.floor(arr.length / 2));
}
