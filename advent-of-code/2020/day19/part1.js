import { buildRegex, parse } from "./helpers.js";

export function part1(puzzleInput) {
  const { rules, messages } = parse(puzzleInput);
  const re = new RegExp(`^${buildRegex(rules, "0")}$`);
  return messages.filter((m) => re.test(m)).length;
}
