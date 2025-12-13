export function solve(puzzleInput: string[]): number {
  const R = puzzleInput.length;
  const C = puzzleInput[0].length;
  let endR = -1;
  let endC = -1;
  const levels: number[][] = [];

  const queue: [number, number, number, number][] = [];

  const dist = new Map<string, number>();
  const key = (r: number, c: number, l: number) => `${r},${c},${l}`;

  const visitedPlatform = new Set<string>();
  const platformKey = (r: number, c: number) => `${r},${c}`;

  for (let r = 0; r < R; r++) {
    levels.push([]);
    for (let c = 0; c < C; c++) {
      const char = puzzleInput[r][c];
      if (char === "S") {
        levels[r].push(0);
        queue.push([r, c, 0, 0]);
        dist.set(key(r, c, 0), 0);
        visitedPlatform.add(platformKey(r, c));
      } else if (char === "E") {
        endR = r;
        endC = c;
        levels[r].push(0);
      } else if (char === "#") {
        levels[r].push(-1);
      } else {
        levels[r].push(parseInt(char, 10));
      }
    }
  }

  const dr = [0, 0, 1, -1];
  const dc = [1, -1, 0, 0];

  while (queue.length > 0) {
    const [r, c, l, t] = queue.shift()!;

    if (r === endR && c === endC && l === 0) {
      return t;
    }

    for (const dl of [1, -1]) {
      const nextL = (l + dl + 10) % 10;
      const nextT = t + 1;
      const nextKey = key(r, c, nextL);

      if (!dist.has(nextKey) || nextT < dist.get(nextKey)!) {
        dist.set(nextKey, nextT);
        queue.push([r, c, nextL, nextT]);
      }
    }

    for (let i = 0; i < 4; i++) {
      const nextR = r + dr[i];
      const nextC = c + dc[i];
      const nextT = t + 1;

      if (nextR >= 0 && nextR < R && nextC >= 0 && nextC < C) {
        const platformLevel = levels[nextR][nextC];

        if (platformLevel === -1) {
          continue;
        }

        const nextPlatformKey = platformKey(nextR, nextC);
        if (visitedPlatform.has(nextPlatformKey)) {
          continue;
        }

        if (platformLevel === l) {
          const nextKey = key(nextR, nextC, l);

          if (!dist.has(nextKey) || nextT < dist.get(nextKey)!) {
            dist.set(nextKey, nextT);
            visitedPlatform.add(nextPlatformKey);
            queue.push([nextR, nextC, l, nextT]);
          }
        }
      }
    }
  }

  return -1;
}
