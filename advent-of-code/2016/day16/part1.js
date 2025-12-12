import { checksum, fillDisk } from "./helpers.js";

export function part1(puzzleInput) {
  const state = puzzleInput[0];
  const disk = fillDisk(state, 272);
  return checksum(disk);
}
