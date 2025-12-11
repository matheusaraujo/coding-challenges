function parseInput(puzzleInput) {
  let columns = [];
  for (const line of puzzleInput) {
    let values = line.split(" ");
    for (let i = 0; i < values.length; i++) {
      columns[i] = columns[i] || [];
      columns[i].push(parseInt(values[i]));
    }
  }
  return columns;
}

module.exports = { parseInput };
