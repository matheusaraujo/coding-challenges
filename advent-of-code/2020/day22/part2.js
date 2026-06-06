import { parse, playRecursive, score } from "./helpers.js";

export function part2(puzzleInput) {
  const [d1, d2] = parse(puzzleInput);
  const [, winnerDeck] = playRecursive(d1, d2);
  return score(winnerDeck);
}
