import { findUniquePath, parseInput } from "./helpers";

export function part2(puzzleInput: string[]): string {
  const tree = parseInput(puzzleInput);
  return findUniquePath(tree)
    .map((n) => n[0])
    .join("");
}
