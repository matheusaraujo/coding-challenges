// sliding-block, this code needs to be refactored
export function part2(puzzleInput) {
  const grid = [];
  let maxX = 0;
  let maxY = 0;
  let empty = null;

  for (const line of puzzleInput) {
    const match = line.match(/node-x(\d+)-y(\d+)\s+\d+T\s+(\d+)T\s+(\d+)T/);
    if (!match) continue;
    const [_, xStr, yStr, usedStr, availStr] = match;
    const x = Number(xStr);
    const y = Number(yStr);
    const used = Number(usedStr);
    const avail = Number(availStr);

    if (!grid[y]) grid[y] = [];
    grid[y][x] = { used, avail };

    if (used === 0) empty = { x, y };
    if (x > maxX) maxX = x;
    if (y > maxY) maxY = y;
  }

  const goal = { x: maxX, y: 0 };

  const walls = new Set();
  for (let y = 0; y <= maxY; y++) {
    for (let x = 0; x <= maxX; x++) {
      if (grid[y][x].used > 100) walls.add(`${x},${y}`);
    }
  }

  function bfs(start, target, walls) {
    const visited = new Set();
    const queue = [{ ...start, steps: 0 }];
    const dirs = [
      { dx: 1, dy: 0 },
      { dx: -1, dy: 0 },
      { dx: 0, dy: 1 },
      { dx: 0, dy: -1 },
    ];

    while (queue.length > 0) {
      const { x, y, steps } = queue.shift();
      const key = `${x},${y}`;
      if (visited.has(key) || walls.has(key)) continue;
      visited.add(key);

      if (x === target.x && y === target.y) return steps;

      for (const d of dirs) {
        const nx = x + d.dx;
        const ny = y + d.dy;
        if (nx >= 0 && nx <= maxX && ny >= 0 && ny <= maxY) {
          queue.push({ x: nx, y: ny, steps: steps + 1 });
        }
      }
    }
  }

  const goalLeft = { x: goal.x - 1, y: goal.y };
  const distEmptyToGoalAdj = bfs(empty, goalLeft, walls);

  const remainingMoves = goal.x - 1;
  const totalSteps = distEmptyToGoalAdj + remainingMoves * 5 + 1;

  return totalSteps;
}
