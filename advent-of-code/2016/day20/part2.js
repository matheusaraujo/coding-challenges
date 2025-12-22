import { parseInput } from "./helpers.js";

export function part2(puzzleInput) {
  const ranges = parseInput(puzzleInput);
  const MAX_IP = 4294967295;
  let allowed = 0;

  allowed += Math.max(0, ranges[0].start);

  for (let i = 1; i < ranges.length; i++) {
    allowed += ranges[i].start - ranges[i - 1].end - 1;
  }

  allowed += Math.max(0, MAX_IP - ranges[ranges.length - 1].end);

  return allowed;
}
