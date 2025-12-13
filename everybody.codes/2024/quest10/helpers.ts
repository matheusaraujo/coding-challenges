export function center(grid: string[][]): string[][] {
  return grid.slice(2, 6).map((row) => row.slice(2, 6));
}

export function runicWord(grid: string[][]): string[][] {
  for (let r = 2; r < 6; r++) {
    for (let c = 2; c < 6; c++) {
      const row = grid[r].filter((a) => a != ".");
      const col = grid.map((row) => row[c]).filter((a) => a != ".");
      const x = row.filter((a) => col.indexOf(a) !== -1);
      if (x && x.length == 1) {
        grid[r][c] = x[0];
      }
    }
  }
  return grid;
}

export function power(grid: string[][]): number {
  if (grid.flat().join("").indexOf(".") != -1) {
    return 0;
  }
  return center(grid)
    .flat()
    .join("")
    .split("")
    .map((v, i) => (i + 1) * (v.charCodeAt(0) - 64))
    .reduce((p, c) => p + c, 0);
}
