export function parseInput(puzzleInput) {
  const columns = [];
  for (const line of puzzleInput) {
    const values = line.split(" ");
    for (let i = 0; i < values.length; i++) {
      columns[i] = columns[i] || [];
      columns[i].push(parseInt(values[i]));
    }
  }
  return columns;
}
