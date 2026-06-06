const DIRS = [[-1, -1], [-1, 0], [-1, 1], [0, -1], [0, 1], [1, -1], [1, 0], [
  1,
  1,
]];

export function simulate(grid, countNeighbors, threshold) {
  let current = grid.map((r) => [...r]);
  while (true) {
    const next = current.map((row, r) =>
      row.map((cell, c) => {
        if (cell === ".") return ".";
        const occ = countNeighbors(current, r, c);
        if (cell === "L" && occ === 0) return "#";
        if (cell === "#" && occ >= threshold) return "L";
        return cell;
      })
    );
    if (next.every((row, r) => row.every((c, ci) => c === current[r][ci]))) {
      return next.flat().filter((c) => c === "#").length;
    }
    current = next;
  }
}

export function adjacentCount(grid, r, c) {
  return DIRS.filter(([dr, dc]) => grid[r + dr]?.[c + dc] === "#").length;
}

export function sightCount(grid, r, c) {
  return DIRS.filter(([dr, dc]) => {
    let nr = r + dr, nc = c + dc;
    while (grid[nr]?.[nc] !== undefined) {
      if (grid[nr][nc] !== ".") return grid[nr][nc] === "#";
      nr += dr;
      nc += dc;
    }
    return false;
  }).length;
}
