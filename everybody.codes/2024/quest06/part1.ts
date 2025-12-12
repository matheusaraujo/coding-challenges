import { findUniquePath, parseInput } from "./helpers";

export function part1(puzzleInput: string[]): any {
  const tree = parseInput(puzzleInput);
  return findUniquePath(tree).join("");
}
