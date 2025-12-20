import { parseInput, solve } from "./helpers";

export function part2(puzzleInput: string[]): string {
  return solve(parseInput(puzzleInput), "Z", 10).toString();
}
