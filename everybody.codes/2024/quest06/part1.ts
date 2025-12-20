import { findUniquePath, parseInput } from "./helpers";

export function part1(puzzleInput: string[]): string {
  const tree = parseInput(puzzleInput);
  return findUniquePath(tree).join("");
}
