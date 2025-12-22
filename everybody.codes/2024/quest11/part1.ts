import { parseInput, solve } from "./helpers";

export function part1(puzzleInput: string[]): any {
  return solve(parseInput(puzzleInput), "A", 4);
}
