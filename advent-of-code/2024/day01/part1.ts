import { parseInput } from "./helpers";

export function part1(puzzleInput: string[]): any {
  const { left, right } = parseInput(puzzleInput);
  return left.reduce(
    (sum: number, l, index) => sum + Math.abs(l - right[index]),
    0,
  );
}
