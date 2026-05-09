export function part2(puzzleInput) {
  return puzzleInput.filter((line) => {
    const [, a, b, char, password] = line.match(
      /^(\d+)-(\d+) (\w): (\w+)$/,
    );
    return (password[Number(a) - 1] === char) !==
      (password[Number(b) - 1] === char);
  }).length;
}
