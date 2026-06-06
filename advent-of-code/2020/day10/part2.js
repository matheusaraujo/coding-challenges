export function part2(puzzleInput) {
  const adapters = puzzleInput.map(Number).sort((a, b) => a - b);
  const chain = [0, ...adapters, adapters.at(-1) + 3];
  const ways = new Array(chain.length).fill(0);
  ways[0] = 1;
  for (let i = 1; i < chain.length; i++) {
    for (let j = i - 1; j >= 0 && chain[i] - chain[j] <= 3; j--) {
      ways[i] += ways[j];
    }
  }
  return ways.at(-1);
}
