import { isTriangle, parseLine } from "./helpers.js";

export function part1(puzzleInput) {
  let validTriangles = 0;
  for (const line of puzzleInput) {
    if (isTriangle(parseLine(line))) {
      validTriangles++;
    }
  }
  return validTriangles;
}
