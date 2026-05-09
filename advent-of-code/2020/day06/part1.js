import { parseGroups } from "./helpers.js";

export function part1(puzzleInput) {
  return parseGroups(puzzleInput)
    .map((group) => new Set(group.join("")).size)
    .reduce((a, b) => a + b, 0);
}
