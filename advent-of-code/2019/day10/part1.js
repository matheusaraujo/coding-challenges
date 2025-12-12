import { bestAsteroid } from "./helpers.js";

export function part1(puzzleInput) {
  return bestAsteroid(puzzleInput).maxVisible;
}
