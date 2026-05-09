export function part1(puzzleInput) {
  return puzzleInput.filter((line) => {
    const [, min, max, char, password] = line.match(
      /^(\d+)-(\d+) (\w): (\w+)$/,
    );
    const count = [...password].filter((c) => c === char).length;
    return count >= Number(min) && count <= Number(max);
  }).length;
}
