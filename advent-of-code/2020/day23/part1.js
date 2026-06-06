import { play } from "./helpers.js";

export function part1(puzzleInput) {
  const cups = puzzleInput[0].split("").map(Number);
  const next = play(cups, 100);
  let result = "", cup = next[1];
  while (cup !== 1) {
    result += cup;
    cup = next[cup];
  }
  return result;
}
