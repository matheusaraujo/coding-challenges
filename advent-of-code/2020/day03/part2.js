import { countTrees } from "./helpers.js";

export function part2(puzzleInput) {
  const slopes = [[1, 1], [3, 1], [5, 1], [7, 1], [1, 2]];
  return slopes.reduce(
    (product, [dx, dy]) => product * countTrees(puzzleInput, dx, dy),
    1,
  );
}
