import { decodeScreen, swapCard } from "./helpers.js";

export function part2(puzzleInput) {
  return decodeScreen(swapCard(puzzleInput));
}
