// Axial hex coordinates: e/w on q axis, ne/sw and nw/se diagonals
const DIRS = {
  e: [1, 0],
  w: [-1, 0],
  ne: [0, 1],
  sw: [0, -1],
  nw: [-1, 1],
  se: [1, -1],
};
const NEIGHBORS = Object.values(DIRS);

export function parseBlack(puzzleInput) {
  const black = new Set();
  for (const line of puzzleInput) {
    let q = 0, r = 0, i = 0;
    while (i < line.length) {
      const dir = line[i] === "e" || line[i] === "w"
        ? line[i++]
        : line.slice(i, i += 2);
      q += DIRS[dir][0];
      r += DIRS[dir][1];
    }
    const key = `${q},${r}`;
    black.has(key) ? black.delete(key) : black.add(key);
  }
  return black;
}

export function step(black) {
  const count = new Map();
  for (const key of black) {
    const [q, r] = key.split(",").map(Number);
    for (const [dq, dr] of NEIGHBORS) {
      const nk = `${q + dq},${r + dr}`;
      count.set(nk, (count.get(nk) || 0) + 1);
    }
  }
  const next = new Set();
  for (const [key, n] of count) {
    if (n === 2 || (n === 1 && black.has(key))) next.add(key);
  }
  return next;
}
