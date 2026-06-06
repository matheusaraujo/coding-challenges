export function simulate(active, dims, cycles) {
  for (let c = 0; c < cycles; c++) {
    const neighbors = new Map();
    for (const key of active) {
      const coord = key.split(",").map(Number);
      for (const delta of deltas(dims)) {
        const nk = coord.map((v, i) => v + delta[i]).join(",");
        neighbors.set(nk, (neighbors.get(nk) ?? 0) + 1);
      }
    }
    const next = new Set();
    for (const [key, count] of neighbors) {
      if (count === 3 || (count === 2 && active.has(key))) next.add(key);
    }
    active = next;
  }
  return active.size;
}

function deltas(dims) {
  let result = [[]];
  for (let d = 0; d < dims; d++) {
    result = result.flatMap((r) => [-1, 0, 1].map((v) => [...r, v]));
  }
  return result.filter((r) => r.some((v) => v !== 0));
}

export function parse(puzzleInput, dims) {
  const active = new Set();
  puzzleInput.forEach((row, y) => {
    [...row].forEach((ch, x) => {
      if (ch === "#") {
        active.add([x, y, ...new Array(dims - 2).fill(0)].join(","));
      }
    });
  });
  return active;
}
