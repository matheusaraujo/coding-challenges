import { parseInput, solve } from "./helpers";

export function part1(puzzleInput: string[]): string {
  return solve(parseInput(puzzleInput)).toString();
}
