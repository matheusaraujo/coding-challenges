import { trees } from "./helpers";

export function part2(puzzleInput: string[]): any {
  return trees(puzzleInput.map((line) => line.split(","))).positions.size;
}
