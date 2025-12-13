import { center, runicWord } from "./helpers";

export function part1(puzzleInput: string[]): any {
  return center(runicWord(puzzleInput.map((line) => line.split(""))))
    .flat()
    .join("");
}
