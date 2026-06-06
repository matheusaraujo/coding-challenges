import { evalP1, tokenize } from "./helpers.js";

export function part1(puzzleInput) {
  return puzzleInput.map((line) => evalP1(tokenize(line), [0])).reduce(
    (a, b) => a + b,
    0,
  );
}
