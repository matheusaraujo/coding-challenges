export function solve(
  dict: Record<string, string[]>,
  initial: string,
  days: number,
): number {
  let counts: Record<string, number> = { [initial]: 1 };

  for (let t = 0; t < days; t++) {
    const newCounts: Record<string, number> = {};
    for (const [key, count] of Object.entries(counts)) {
      if (dict[key]) {
        for (const child of dict[key]) {
          newCounts[child] = (newCounts[child] || 0) + count;
        }
      }
    }
    counts = newCounts;
  }

  return Object.values(counts).reduce((a, b) => a + b, 0);
}

export function parseInput(puzzleInput: string[]): Record<string, string[]> {
  const dict: Record<string, string[]> = {};
  for (const line of puzzleInput) {
    const [id, gen] = line.split(":");
    dict[id] = gen.split(",");
  }
  return dict;
}
