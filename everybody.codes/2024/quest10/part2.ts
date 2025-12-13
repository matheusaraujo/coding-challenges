import { power, runicWord } from "./helpers";

export function part2(puzzleInput: string[]): any {
  let result = 0;
  for (let r = 0; r < puzzleInput.length; r += 9) {
    for (let c = 0; c < puzzleInput[0].length; c += 9) {
      const grid = puzzleInput
        .slice(r, r + 8)
        .map((a) => a.slice(c, c + 8))
        .map((a) => a.split(""));
      result += power(runicWord(grid));
    }
  }
  return result;
}
