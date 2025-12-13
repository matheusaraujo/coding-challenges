import { parseInput, solve } from "./helpers";

export function part2(puzzleInput: string[]): any {
  return solve(parseInput(puzzleInput), "Z", 10);
}
