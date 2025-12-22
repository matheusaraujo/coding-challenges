// based on https://www.reddit.com/r/adventofcode/comments/5j4lp1/comment/dbdihvu/
export function part2(puzzleInput) {
  const N = parseInt(puzzleInput[0]);

  let pow = 1;
  while (pow * 3 < N) {
    pow *= 3;
  }

  if (N === pow) {
    return N;
  } else if (N <= 2 * pow) {
    return N - pow;
  } else {
    return 2 * N - 3 * pow;
  }
}
