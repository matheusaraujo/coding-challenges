import { trees } from "./helpers";

export function part1(puzzleInput: string[]): string {
  return trees(puzzleInput.map((line) => line.split(","))).height.toString();
}
