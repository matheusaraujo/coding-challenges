import { parse } from "./helpers.js";

export function part1(puzzleInput) {
  const { fields, nearby } = parse(puzzleInput);
  const validators = Object.values(fields);
  return nearby.flat().filter((v) => !validators.some((fn) => fn(v))).reduce(
    (a, b) => a + b,
    0,
  );
}
