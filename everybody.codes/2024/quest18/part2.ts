import { bfs, parseInput } from "./helpers";

export function part2(puzzleInput: string[]): any {
  const { starts, palms, grid } = parseInput(puzzleInput);
  return bfs(starts, palms, grid).maxTime;
}
