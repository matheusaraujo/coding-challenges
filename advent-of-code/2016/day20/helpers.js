export function parseInput(input) {
  const ranges = input.map((line) => {
    const [start, end] = line.split("-").map(Number);
    return { start, end };
  });

  ranges.sort((a, b) => a.start - b.start);

  const merged = [];
  for (const range of ranges) {
    if (merged.length === 0) {
      merged.push({ ...range });
    } else {
      const last = merged[merged.length - 1];
      if (range.start <= last.end + 1) {
        last.end = Math.max(last.end, range.end);
      } else {
        merged.push({ ...range });
      }
    }
  }

  return merged;
}
