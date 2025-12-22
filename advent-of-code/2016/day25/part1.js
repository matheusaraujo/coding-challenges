import { Computer } from "./helpers.js";

export function part1(puzzleInput) {
  let a = 1;
  while (true) {
    const comp = new Computer(puzzleInput, { a: a });
    if (comp.execute(10) === true) {
      return a;
    }
    a++;
  }
}
