import { play } from "./helpers.js";

export function part2(puzzleInput) {
  const initial = puzzleInput[0].split("").map(Number);
  const cups = [
    ...initial,
    ...Array.from({ length: 1_000_000 - 9 }, (_, i) => i + 10),
  ];
  const next = play(cups, 10_000_000);
  return next[1] * next[next[1]];
}
