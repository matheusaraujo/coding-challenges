import { Grid } from "./helpers";
import { solvePart2 } from "./part2";

// This is a heuristic, not a fully general solution.
// It works based on the assumpions for the map layout:
// the middle section only connects to left/right at the bottom, and herbs are close by.
// We slightly tweak the input (rename one herb: K → L) to force both herbs to be visited.
// A proper generic solution would need BFS over a growing state graph — which is overkill here.
export function part3(puzzleInput: string[]): any {
  const map: Grid = puzzleInput.map((line: string) => line.split(""));
  const h = map.length;
  const w = 85;

  const maps: Grid[] = [];
  for (let i = 0; i < 3; i++) {
    maps[i] = map.map((row) => row.slice(w * i, w * (i + 1)));
  }

  maps[1][h - 2][maps[1][h - 2].indexOf("K")] = "L";

  const map0Dist = solvePart2(maps[0], [w - 1, h - 2]) + 1;
  const map2Dist = solvePart2(maps[2], [0, h - 2]) + 1;

  const midRow = maps[1][h - 2];
  const dist =
    solvePart2(maps[1]) +
    2 +
    2 * midRow.indexOf("L") +
    2 * (w - 1 - midRow.indexOf("K"));

  return map0Dist + map2Dist + dist;
}
