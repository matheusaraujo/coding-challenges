import { seatId } from "./helpers.js";

export function part2(puzzleInput) {
  const ids = puzzleInput.map(seatId).sort((a, b) => a - b);
  for (let i = 1; i < ids.length; i++) {
    if (ids[i] !== ids[i - 1] + 1) return ids[i] - 1;
  }
  return null;
}
