import { bfs, parseInput } from "./helpers";

export function part1(puzzleInput: string[]): string {
  const { starts, palms, grid } = parseInput(puzzleInput);
  return bfs(starts, palms, grid).maxTime.toString();
}
