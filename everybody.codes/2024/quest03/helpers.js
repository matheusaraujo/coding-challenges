function solve(grid, dirs) {
  grid = border(grid);
  const h = grid.length;
  const w = grid[0].length;

  const dist = Array.from({ length: h }, () => Array(w).fill(-1));
  const q = [];
  let qHead = 0;

  let result = 0;

  for (let i = 0; i < h; i++) {
    for (let j = 0; j < w; j++) {
      if (grid[i][j] !== "#") {
        dist[i][j] = 0;
        q.push([i, j]);
      }
    }
  }

  while (qHead < q.length) {
    const [r, c] = q[qHead++];
    for (const [dr, dc] of dirs) {
      const nr = r + dr,
        nc = c + dc;
      if (nr >= 0 && nr < h && nc >= 0 && nc < w && dist[nr][nc] === -1) {
        dist[nr][nc] = dist[r][c] + 1;
        q.push([nr, nc]);
        result += dist[nr][nc];
      }
    }
  }

  return result;
}

function border(inputGrid) {
  const originalH = inputGrid.length;
  const originalW = inputGrid[0].length;
  const grid = [];

  grid.push(Array(originalW + 2).fill("."));

  for (let i = 0; i < originalH; i++) {
    const row = inputGrid[i];
    const rowChars = typeof row === "string" ? row.split("") : row;
    grid.push([".", ...rowChars, "."]);
  }

  grid.push(Array(originalW + 2).fill("."));

  return grid;
}

module.exports = solve;
