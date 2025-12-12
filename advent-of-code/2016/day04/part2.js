import { Room } from "./helpers.js";

export function part2(puzzleInput) {
  for (let i = 0; i < puzzleInput.length; i++) {
    const room = new Room(puzzleInput[i]);
    if (room.decrypt() === "northpole object storage") {
      return room.sectorId;
    }
  }
}
