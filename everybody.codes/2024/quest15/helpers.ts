export type Point = [number, number];
export type Grid = string[][];
export type DistanceMap = number[][];

const addVect = (a: Point, b: Point): Point => [a[0] + b[0], a[1] + b[1]];

const keyOf = (v: Point): string => v.join("_");

const MOVES: Record<string, Point> = {
  U: [0, 1],
  D: [0, -1],
  R: [1, 0],
  L: [-1, 0],
};

const moveVectors: Point[] = Object.values(MOVES);

export type DistanceCache = Record<string, DistanceMap>;

export function distanceMap(
  map: Grid,
  from: Point,
  cache: DistanceCache,
): DistanceMap {
  const cacheKey = keyOf(from);
  if (cache[cacheKey]) return cache[cacheKey];

  const stack: { p: Point; dist: number }[] = [
    {
      p: [...from],
      dist: 0,
    },
  ];

  const filled: DistanceMap = map.map((row) => row.map(() => Infinity));

  while (stack.length !== 0) {
    const cur = stack.shift()!;
    const [x, y] = cur.p;

    if (filled[y][x] <= cur.dist) continue;
    filled[y][x] = cur.dist;

    for (const m of moveVectors) {
      const np = addVect(cur.p, m);
      const [nx, ny] = np;

      if (nx < 0 || nx >= map[0].length || ny < 0 || ny >= map.length) continue;

      if (["~", "#"].includes(map[ny][nx])) continue;
      if (filled[ny][nx] <= cur.dist + 1) continue;

      stack.push({ p: np, dist: cur.dist + 1 });
    }
  }

  cache[cacheKey] = filled;
  return filled;
}

export function permutator<T>(input: T[]): T[][] {
  const result: T[][] = [];

  const permute = (arr: T[], m: T[] = []) => {
    if (arr.length === 0) {
      result.push(m);
    } else {
      for (let i = 0; i < arr.length; i++) {
        const curr = arr.slice();
        const next = curr.splice(i, 1);
        permute(curr, m.concat(next));
      }
    }
  };

  permute(input);
  return result;
}
