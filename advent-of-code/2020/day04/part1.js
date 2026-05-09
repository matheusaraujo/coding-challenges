import { hasRequiredFields, parsePassports } from "./helpers.js";

export function part1(puzzleInput) {
  return parsePassports(puzzleInput).filter(hasRequiredFields).length;
}
