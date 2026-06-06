import { evalP2, tokenize } from "./helpers.js";

export function part2(puzzleInput) {
  return puzzleInput.map((line) => evalP2(tokenize(line), [0])).reduce(
    (a, b) => a + b,
    0,
  );
}
