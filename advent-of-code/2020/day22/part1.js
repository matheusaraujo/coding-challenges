import { parse, score } from "./helpers.js";

export function part1(puzzleInput) {
  const [d1, d2] = parse(puzzleInput);
  while (d1.length && d2.length) {
    const [a, b] = [d1.shift(), d2.shift()];
    a > b ? d1.push(a, b) : d2.push(b, a);
  }
  return score(d1.length ? d1 : d2);
}
