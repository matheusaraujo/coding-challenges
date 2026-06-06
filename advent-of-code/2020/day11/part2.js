import { sightCount, simulate } from "./helpers.js";

export function part2(puzzleInput) {
  return simulate(puzzleInput.map((r) => r.split("")), sightCount, 5);
}
