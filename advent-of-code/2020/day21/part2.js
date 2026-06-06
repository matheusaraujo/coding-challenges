import { parse, resolve } from "./helpers.js";

export function part2(puzzleInput) {
  const assigned = resolve(parse(puzzleInput));
  return [...assigned.keys()].sort().map((a) => assigned.get(a)).join(",");
}
