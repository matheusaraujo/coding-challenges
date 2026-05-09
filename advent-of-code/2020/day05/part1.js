import { seatId } from "./helpers.js";

export function part1(puzzleInput) {
  return Math.max(...puzzleInput.map(seatId));
}
