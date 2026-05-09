export function countTrees(grid, dx, dy) {
  const width = grid[0].length;
  let trees = 0;
  for (let x = 0, y = 0; y < grid.length; x += dx, y += dy) {
    if (grid[y][x % width] === "#") trees++;
  }
  return trees;
}
