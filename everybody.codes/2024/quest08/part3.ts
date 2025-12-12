export function part3(puzzleInput: string[]): any {
  return solve(parseInt(puzzleInput[0]), 10, 202400000);
}

const solve = (hp = 2, hpa = 5, blocks = 160) => {
  const nextThickness = (t: number) => ((t * hp) % hpa) + hpa;
  const toBeRemoved = (h: number) => (hp * w * h) % hpa;
  const heights = [1];
  let used = 1,
    w = 1,
    t = 1;

  while (used < blocks) {
    t = nextThickness(t);
    for (let i = 0; i <= (w + 1) / 2; i++) {
      if (heights[i] === undefined) heights[i] = 0;
      heights[i] += t;
    }
    used += (w + 2) * t;
    w += 2;
  }

  for (let i = 0; i < heights.length; i++) {
    const upperFrameBlocks = heights[i] - (heights[i + 1] || 0);
    const removables = heights[i] - upperFrameBlocks;
    if (removables == 0) continue;
    let removed = Math.min(toBeRemoved(heights[i]), removables);
    if (i > 0) removed = removed * 2;
    used -= removed;
  }

  return used - blocks;
};
