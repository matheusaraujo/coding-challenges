import { parseInstructions, processBots } from "./helpers.js";

export function part1(puzzleInput) {
  const { botChips, botRules } = parseInstructions(puzzleInput);
  const { targetBot } = processBots(botChips, botRules, [17, 61]);
  return targetBot;
}
