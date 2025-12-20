import { dp } from "./helpers";

export function part1(puzzleInput: string[]): string {
  return puzzleInput
    .map((v: string) => parseInt(v))
    .map((i: number) => {
      return dp(i, [10, 5, 3, 1])[i];
    })
    .reduce((acc, n) => acc + n, 0)
    .toString();
}
