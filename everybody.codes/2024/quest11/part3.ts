import { parseInput, solve } from "./helpers";

export function part3(puzzleInput: string[]): string {
  const dict = parseInput(puzzleInput);
  let min = Number.MAX_SAFE_INTEGER,
    max = Number.MIN_SAFE_INTEGER;

  for (const [k] of Object.entries(dict)) {
    const x = solve(dict, k, 20);
    min = Math.min(min, x);
    max = Math.max(max, x);
  }

  return (max - min).toString();
}
