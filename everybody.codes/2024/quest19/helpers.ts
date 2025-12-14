export type Coord = [number, number];
export type Grid<T> = T[][];

export const offsets: Coord[] = [
  [-1, -1],
  [-1, 0],
  [-1, 1],
  [0, 1],
  [1, 1],
  [1, 0],
  [1, -1],
  [0, -1],
];

export function parseInput(puzzleInput: string[]): {
  sequence: string[];
  grid: Grid<string>;
} {
  const sequence = puzzleInput[0].split("");
  const grid = puzzleInput.slice(2).map((line) => line.split(""));
  return { sequence, grid };
}

export function findAnswer(message: string): string {
  return message.slice(message.indexOf(">") + 1, message.indexOf("<"));
}

export function solve(
  grid: Grid<string>,
  sequence: string[],
  rounds: number,
): string {
  const cycles = findCycles(grid, sequence);
  const height = grid.length;
  const width = grid[0].length;

  const newgrid: Grid<string> = Array.from({ length: height }, () =>
    Array(width).fill(""),
  );

  for (const cycle of cycles) {
    const len = cycle.length;
    for (let i = 0; i < len; i++) {
      const [sy, sx] = cycle[i];
      const [dy, dx] = cycle[(i + rounds) % len];
      newgrid[dy][dx] = grid[sy][sx];
    }
  }

  return newgrid.map((row) => row.join("")).join("\n");
}

function findCycles<T>(grid: Grid<T>, sequence: string[]): Coord[][] {
  const height = grid.length;
  const width = grid[0].length;

  const tgrid: Grid<Coord> = Array.from({ length: height }, (_, y) =>
    Array.from({ length: width }, (_, x) => [y, x]),
  );

  for (let y = 1, i = 0; y < height - 1; y++) {
    for (let x = 1; x < width - 1; x++, i++) {
      const dir = sequence[i % sequence.length];
      const dj = dir === "R" ? 1 : -1;
      const vals = offsets.map(([dy, dx]) => tgrid[y + dy][x + dx]);
      vals.forEach((val, j) => {
        const [dy, dx] = offsets[(j + dj + 8) % 8];
        tgrid[y + dy][x + dx] = val;
      });
    }
  }

  const transition = new Map<string, Coord>();
  for (let y = 0; y < height; y++) {
    for (let x = 0; x < width; x++) {
      const [sy, sx] = tgrid[y][x];
      transition.set(`${sy},${sx}`, [y, x]);
    }
  }

  const cycles: Coord[][] = [];
  const seen = new Set<string>();

  for (let sy = 0; sy < height; sy++) {
    for (let sx = 0; sx < width; sx++) {
      const startKey = `${sy},${sx}`;
      if (seen.has(startKey)) continue;

      const cycle: Coord[] = [];
      let y = sy;
      let x = sx;

      while (!seen.has(`${y},${x}`)) {
        cycle.push([y, x]);
        seen.add(`${y},${x}`);
        const next = transition.get(`${y},${x}`)!;
        [y, x] = next;
      }

      cycles.push(cycle);
    }
  }

  return cycles;
}
