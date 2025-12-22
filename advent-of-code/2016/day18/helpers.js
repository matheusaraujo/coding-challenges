export function solve(puzzleInput, totalRows) {
  let row = puzzleInput[0].split("").map((c) => c === "^");
  const width = row.length;
  let safeCount = row.filter((tile) => !tile).length;

  for (let r = 1; r < totalRows; r++) {
    const newRow = new Array(width);

    for (let i = 0; i < width; i++) {
      const left = i > 0 ? row[i - 1] : false;
      const center = row[i];
      const right = i < width - 1 ? row[i + 1] : false;

      newRow[i] = (left && center && !right) ||
        (!left && center && right) ||
        (left && !center && !right) ||
        (!left && !center && right);
    }

    safeCount += newRow.filter((tile) => !tile).length;
    row = newRow;
  }

  return safeCount;
}
