import { parseInput, solve } from "./helpers";

export function part1(puzzleInput: string[]): number {
  return solve(parseInput(puzzleInput), "A", 4);
}
