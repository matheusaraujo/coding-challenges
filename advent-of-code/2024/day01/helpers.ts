export const parseInput = (
  puzzleInput: string[],
): { left: number[]; right: number[] } => {
  const left: number[] = [];
  const right: number[] = [];

  puzzleInput.forEach((line) => {
    const parts = line.split("   ");
    left.push(parseInt(parts[0], 10));
    right.push(parseInt(parts[1], 10));
  });

  left.sort((a, b) => a - b);
  right.sort((a, b) => a - b);

  return { left, right };
};
