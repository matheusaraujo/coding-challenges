import { parseGroups } from "./helpers.js";

export function part2(puzzleInput) {
  return parseGroups(puzzleInput)
    .map((group) =>
      group
        .map((person) => new Set(person))
        .reduce((acc, set) => new Set([...acc].filter((c) => set.has(c))))
        .size
    )
    .reduce((a, b) => a + b, 0);
}
