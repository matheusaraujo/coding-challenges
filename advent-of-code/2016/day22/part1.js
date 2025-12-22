export function part1(puzzleInput) {
  // Parse nodes
  const nodes = [];
  for (const line of puzzleInput) {
    const match = line.match(/node-x(\d+)-y(\d+)\s+\d+T\s+(\d+)T\s+(\d+)T/);
    if (match) {
      const [, x, y, used, avail] = match;
      nodes.push({
        x: Number(x),
        y: Number(y),
        used: Number(used),
        avail: Number(avail),
      });
    }
  }

  let viablePairs = 0;

  // Check all pairs (A,B)
  for (let i = 0; i < nodes.length; i++) {
    const A = nodes[i];
    if (A.used === 0) continue;

    for (let j = 0; j < nodes.length; j++) {
      if (i === j) continue;
      const B = nodes[j];
      if (A.used <= B.avail) viablePairs++;
    }
  }

  return viablePairs;
}
