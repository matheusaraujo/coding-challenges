import { bfs, parseInput } from "./helpers";

export function part3(puzzleInput: string[]): string {
  const { palms, grid, wells } = parseInput(puzzleInput);
  let minTotalTime = Infinity;

  for (const well of wells) {
    minTotalTime = Math.min(minTotalTime, bfs([well], palms, grid).totalTime);
  }

  return minTotalTime.toString();
}
