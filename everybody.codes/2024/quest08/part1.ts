export function part1(puzzleInput: string[]): string {
  return solve(parseInt(puzzleInput[0])).toString();
}

function solve(n: number): number {
  let acc = 1,
    b = 1;

  while (acc < n) {
    b += 2;
    acc += b;
  }

  return (acc - n) * b;
}
