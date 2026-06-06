import { buildRegex, parse } from "./helpers.js";

export function part2(puzzleInput) {
  const { rules, messages } = parse(puzzleInput);
  const r42 = buildRegex(rules, "42");
  const r31 = buildRegex(rules, "31");
  // rule 8: 42+
  // rule 11: 42{n} 31{n} for n >= 1, bounded by max message length
  const maxN = Math.ceil(messages[0].length / 2);
  const r11 = Array.from(
    { length: maxN },
    (_, i) => `(?:${r42}){${i + 1}}(?:${r31}){${i + 1}}`,
  ).join("|");
  const re = new RegExp(`^(?:${r42})+(?:${r11})$`);
  return messages.filter((m) => re.test(m)).length;
}
