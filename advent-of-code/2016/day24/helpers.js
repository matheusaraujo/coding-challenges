export function solve(puzzleInput, returnToZero = false) {
  const grid = puzzleInput.map((line) => line.split(""));
  const points = {};

  for (let y = 0; y < grid.length; y++) {
    for (let x = 0; x < grid[y].length; x++) {
      if (/\d/.test(grid[y][x])) {
        points[grid[y][x]] = { x, y };
      }
    }
  }

  const pointKeys = Object.keys(points).sort();
  const distances = {};

  for (let i = 0; i < pointKeys.length; i++) {
    for (let j = i + 1; j < pointKeys.length; j++) {
      const dist = bfs(grid, points[pointKeys[i]], points[pointKeys[j]]);
      distances[`${pointKeys[i]}-${pointKeys[j]}`] = dist;
      distances[`${pointKeys[j]}-${pointKeys[i]}`] = dist;
    }
  }

  const targets = pointKeys.filter((k) => k !== "0");
  const routes = permutations(targets);

  let minDistance = Infinity;

  for (const route of routes) {
    let totalDist = 0;
    let currentPos = "0";

    for (const nextPos of route) {
      totalDist += distances[`${currentPos}-${nextPos}`];
      currentPos = nextPos;
    }

    if (returnToZero) {
      totalDist += distances[`${currentPos}-0`];
    }

    minDistance = Math.min(minDistance, totalDist);
  }

  return minDistance;
}

function bfs(grid, start, end) {
  const queue = [{ x: start.x, y: start.y, dist: 0 }];
  const visited = new Set([`${start.x},${start.y}`]);
  const dirs = [[0, 1], [0, -1], [1, 0], [-1, 0]];

  while (queue.length > 0) {
    const { x, y, dist } = queue.shift();
    if (x === end.x && y === end.y) return dist;

    for (const [dx, dy] of dirs) {
      const nx = x + dx, ny = y + dy;
      const key = `${nx},${ny}`;
      if (grid[ny] && grid[ny][nx] !== "#" && !visited.has(key)) {
        visited.add(key);
        queue.push({ x: nx, y: ny, dist: dist + 1 });
      }
    }
  }
}

function permutations(arr) {
  if (arr.length <= 1) return [arr];
  const perms = [];
  for (let i = 0; i < arr.length; i++) {
    const char = arr[i];
    const remaining = [...arr.slice(0, i), ...arr.slice(i + 1)];
    for (const p of permutations(remaining)) {
      perms.push([char, ...p]);
    }
  }
  return perms;
}
