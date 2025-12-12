import { Room } from "./helpers.js";

export function part1(puzzleInput) {
  let result = 0;
  for (const line of puzzleInput) {
    const room = new Room(line);
    if (room.isReal()) {
      result += room.sectorId;
    }
  }
  return result;
}
