import { trees } from "./helpers";

export function part2(puzzleInput: string[]): string {
  return trees(
    puzzleInput.map((line) => line.split(",")),
  ).positions.size.toString();
}
