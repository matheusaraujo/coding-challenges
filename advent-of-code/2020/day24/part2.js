import { parseBlack, step } from "./helpers.js";

export function part2(puzzleInput) {
  let black = parseBlack(puzzleInput);
  for (let i = 0; i < 100; i++) black = step(black);
  return black.size;
}
