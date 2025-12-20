import { DistanceCache, distanceMap, Grid, Point } from "./helpers";

export function part1(puzzleInput: string[]): string {
  const map: Grid = puzzleInput.map((line) => line.split(""));
  const cache: DistanceCache = {};

  let start: Point = [0, 0];
  const fruits: { type: string; p: Point }[] = [];

  map.forEach((row: string[], y: number) =>
    row.forEach((v: string, x: number) => {
      if (y === 0 && v === ".") start = [x, y];
      if (["~", "#", "."].includes(v)) return;
      fruits.push({ type: v, p: [x, y] });
    }),
  );

  const dmap = distanceMap(map, start, cache);

  return (2 * Math.min(...fruits.map((f) => dmap[f.p[1]][f.p[0]]))).toString();
}
