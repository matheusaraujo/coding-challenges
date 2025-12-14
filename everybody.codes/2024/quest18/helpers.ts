interface Coord {
  r: number;
  c: number;
}

const directions = [
  { dr: -1, dc: 0 },
  { dr: 1, dc: 0 },
  { dr: 0, dc: -1 },
  { dr: 0, dc: 1 },
];

export function parseInput(puzzleInput: string[]): {
  starts: Coord[];
  palms: Coord[];
  grid: string[][];
  wells: Coord[];
} {
  const grid = puzzleInput.map((line) => line.split(""));
  const rows = grid.length;
  const cols = grid[0].length;
  const starts: Coord[] = [];
  const palms: Coord[] = [];
  const wells: Coord[] = [];

  for (let r = 0; r < rows; r++) {
    for (let c = 0; c < cols; c++) {
      if (grid[r][c] === "P") palms.push({ r, c });

      if (grid[r][c] === ".") {
        wells.push({ r, c });
        if (c === 0) {
          starts.push({ r, c });
        } else if (c === cols - 1) {
          starts.push({ r, c });
        }
      }
    }
  }

  return { starts, palms, grid, wells };
}

export function bfs(
  starts: Coord[],
  palms: Coord[],
  grid: string[][],
): { maxTime: number; totalTime: number } {
  const rows = grid.length;
  const cols = grid[0].length;
  const queue: [Coord, number][] = [];

  const visitedTimes: number[][] = Array.from({ length: rows }, () =>
    new Array(cols).fill(Infinity),
  );

  for (const start of starts) {
    if (grid[start.r][start.c] === "." || grid[start.r][start.c] === "P") {
      queue.push([start, 0]);
      visitedTimes[start.r][start.c] = 0;
    }
  }

  const palmKeys = new Set(palms.map((p) => `${p.r},${p.c}`));
  const palmTimes = new Map<string, number>();
  let maxTime = 0;

  while (queue.length > 0) {
    const [currentCoord, currentTime] = queue.shift()!;
    const { r, c } = currentCoord;

    const key = `${r},${c}`;
    if (palmKeys.has(key) && !palmTimes.has(key)) {
      palmTimes.set(key, currentTime);
      maxTime = Math.max(maxTime, currentTime);

      if (palmTimes.size === palms.length) {
        break;
      }
    }

    for (const dir of directions) {
      const nr = r + dir.dr;
      const nc = c + dir.dc;
      const newTime = currentTime + 1;

      if (nr < 0 || nr >= rows || nc < 0 || nc >= cols) {
        continue;
      }

      const cell = grid[nr][nc];
      const isTraversable = cell === "." || cell === "P";

      if (isTraversable && newTime < visitedTimes[nr][nc]) {
        visitedTimes[nr][nc] = newTime;
        queue.push([{ r: nr, c: nc }, newTime]);
      }
    }
  }

  const totalTime = Array.from(palmTimes.values()).reduce((a, b) => a + b, 0);
  return { maxTime, totalTime };
}
