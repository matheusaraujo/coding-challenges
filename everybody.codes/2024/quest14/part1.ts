import { trees } from "./helpers";

export function part1(puzzleInput: string[]): any {
  return trees(puzzleInput.map((line) => line.split(","))).height;
}
