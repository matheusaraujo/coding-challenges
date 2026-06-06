import { adjacentCount, simulate } from "./helpers.js";

export function part1(puzzleInput) {
  return simulate(puzzleInput.map((r) => r.split("")), adjacentCount, 4);
}
