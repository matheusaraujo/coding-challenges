import { findAnswer, parseInput, solve } from "./helpers";

export function part3(puzzleInput: string[]): any {
  const { sequence, grid } = parseInput(puzzleInput);
  return findAnswer(solve(grid, sequence, 1_048_576_000));
}
