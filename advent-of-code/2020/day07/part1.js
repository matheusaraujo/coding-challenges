import { parseRules } from "./helpers.js";

export function part1(puzzleInput) {
  const rules = parseRules(puzzleInput);
  const canContainGold = new Set();

  const canHoldGold = (color) => {
    if (canContainGold.has(color)) return true;
    return (rules[color] || []).some(
      ({ color: inner }) => inner === "shiny gold" || canHoldGold(inner),
    );
  };

  for (const color of Object.keys(rules)) {
    if (canHoldGold(color)) canContainGold.add(color);
  }

  return canContainGold.size;
}
