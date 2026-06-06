import { parse, resolve } from "./helpers.js";

export function part1(puzzleInput) {
  const foods = parse(puzzleInput);
  const dangerous = new Set(resolve(foods).values());
  return foods.flatMap(({ ingredients }) => ingredients).filter(
    (ing) => !dangerous.has(ing),
  ).length;
}
