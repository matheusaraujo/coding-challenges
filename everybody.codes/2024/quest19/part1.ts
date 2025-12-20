import { offsets, parseInput } from "./helpers";

export function part1(puzzleInput: string[]): string {
  const { sequence, grid } = parseInput(puzzleInput);
  const height = grid.length;
  const width = grid[0].length;

  for (let y = 1, i = 0; y < height - 1; y++) {
    for (let x = 1; x < width - 1; x++, i++) {
      const dir = sequence[i % sequence.length];
      const dj = dir === "R" ? 1 : -1;
      const vals = offsets.map(([dy, dx]) => grid[y + dy][x + dx]);
      vals.forEach((val, j) => {
        const [dy, dx] = offsets[(j + dj + 8) % 8];
        grid[y + dy][x + dx] = val;
      });
    }
  }

  return grid[1].slice(1, -1).join("");
}
