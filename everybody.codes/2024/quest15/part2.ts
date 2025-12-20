import {
  DistanceCache,
  DistanceMap,
  distanceMap,
  Grid,
  permutator,
  Point,
} from "./helpers";

export function part2(puzzleInput: string[]): string {
  return solvePart2(
    puzzleInput.map((line: string) => line.split("")),
  ).toString();
}

export function solvePart2(map: Grid, start?: Point, end?: Point): number {
  const cache: DistanceCache = {};
  const fruits: { type: string; p: Point }[] = [];
  const fruitTypes: Record<string, number> = {};

  map.forEach((row, y) =>
    row.forEach((v, x) => {
      if (y === 0 && v === "." && !start) start = [x, y];
      if (["~", "#", "."].includes(v)) return;

      fruits.push({ type: v, p: [x, y] });
      fruitTypes[v] = (fruitTypes[v] ?? 0) + 1;
    }),
  );

  start = start!;
  end ??= [...start];

  let min = Infinity;
  const dmap0 = distanceMap(map, start, cache);

  const recur = (
    seq: string[],
    lastDist: number,
    lastDmap: DistanceMap,
  ): void => {
    fruits
      .filter((f) => f.type === seq[0])
      .forEach((fruit) => {
        let dist = lastDist + lastDmap[fruit.p[1]][fruit.p[0]];
        if (dist >= min) return;

        const dmap = distanceMap(map, fruit.p, cache);

        if (seq.length > 1) {
          recur(seq.slice(1), dist, dmap);
        } else {
          dist += dmap[end![1]][end![0]];
          if (dist < min) min = dist;
        }
      });
  };

  permutator(Object.keys(fruitTypes)).forEach((seq) => recur(seq, 0, dmap0));

  return min;
}
