import { dp } from "./helpers";

export function part2(puzzleInput: string[]): string {
  return puzzleInput
    .map((v: string) => parseInt(v))
    .map((i: number) => {
      return dp(i, [30, 25, 24, 20, 16, 15, 10, 5, 3, 1])[i];
    })
    .reduce((acc, n) => acc + n, 0)
    .toString();
}
