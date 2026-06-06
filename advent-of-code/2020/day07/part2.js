import { parseRules } from "./helpers.js";

export function part2(puzzleInput) {
  const rules = parseRules(puzzleInput);

  const countBags = (color) =>
    (rules[color] || []).reduce(
      (sum, { count, color: inner }) => sum + count + count * countBags(inner),
      0,
    );

  return countBags("shiny gold");
}
