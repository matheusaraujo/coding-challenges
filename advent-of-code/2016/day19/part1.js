// https://en.wikipedia.org/wiki/Josephus_problem
export function part1(puzzleInput) {
  const N = puzzleInput[0];
  let power = 1;
  while (power * 2 <= N) power *= 2;
  const L = N - power;
  return 2 * L + 1;
}
