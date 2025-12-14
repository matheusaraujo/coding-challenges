import { findAnswer, parseInput, solve } from "./helpers";

export function part2(puzzleInput: string[]): any {
  const { sequence, grid } = parseInput(puzzleInput);
  return findAnswer(solve(grid, sequence, 100));
}
