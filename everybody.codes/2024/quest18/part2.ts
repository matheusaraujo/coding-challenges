import { bfs, parseInput } from "./helpers";

export function part2(puzzleInput: string[]): number {
  const { starts, palms, grid } = parseInput(puzzleInput);
  return bfs(starts, palms, grid).maxTime;
}
