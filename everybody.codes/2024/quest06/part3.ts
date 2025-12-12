import { findUniquePath, parseInput } from "./helpers";

export function part3(puzzleInput: string[]): any {
  const tree = parseInput(puzzleInput);
  return findUniquePath(tree)
    .map((n) => n[0])
    .join("");
}
