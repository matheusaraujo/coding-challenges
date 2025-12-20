import { findAnswer, parseInput, solve } from "./helpers";

export function part2(puzzleInput: string[]): string {
  const { sequence, grid } = parseInput(puzzleInput);
  return findAnswer(solve(grid, sequence, 100));
}
