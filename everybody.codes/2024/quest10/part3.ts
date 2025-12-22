import { power, runicWord } from "./helpers";

export function part3(puzzleInput: string[]): any {
  let result = 0,
    prev = -1;
  const gridInput = puzzleInput.map((row) => row.split(""));

  while (result != prev) {
    prev = result;
    result = 0;
    for (let r = 0; r < gridInput.length - 2; r += 6) {
      for (let c = 0; c < gridInput[0].length - 2; c += 6) {
        let grid = gridInput.slice(r, r + 8).map((row) => row.slice(c, c + 8));
        grid = runicWord(grid);
        grid = replace(grid);
        const p = power(grid);
        if (p > 0) {
          for (let i = 0; i < 8; i++) {
            for (let j = 0; j < 8; j++) {
              gridInput[r + i][c + j] = grid[i][j];
            }
          }
        }
        result += p;
      }
    }
  }

  return result;
}

function replace(grid: string[][]) {
  const rows = grid.length;
  const cols = grid[0].length;

  for (let r = 2; r < rows - 2; r++) {
    for (let c = 2; c < cols - 2; c++) {
      if (grid[r][c] === ".") {
        const ch = findUnique(grid, r, c);
        if (ch) {
          grid[r][c] = ch;
          for (let x = 0; x < 8; x++) {
            if (grid[r][x] === "?") grid[r][x] = ch;
            if (grid[x][c] === "?") grid[x][c] = ch;
          }
        }
      }
    }
  }

  return grid;
}

function findUnique(grid: string[][], r: number, c: number): string | null {
  const row = grid[r];
  const col = grid.map((r) => r[c]);

  const freq = new Map<string, number>();

  for (const ch of [...row, ...col]) {
    if (ch !== "." && ch !== "?") {
      freq.set(ch, (freq.get(ch) ?? 0) + 1);
    }
  }

  const matches: string[] = [];

  for (const [char, count] of freq) {
    if (count == 1) {
      matches.push(char);
    }
  }

  return matches.length === 1 ? matches[0] : null;
}
