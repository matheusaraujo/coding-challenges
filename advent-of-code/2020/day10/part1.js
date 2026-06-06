export function part1(puzzleInput) {
  const adapters = puzzleInput.map(Number).sort((a, b) => a - b);
  const chain = [0, ...adapters, adapters.at(-1) + 3];
  let ones = 0, threes = 0;
  for (let i = 1; i < chain.length; i++) {
    const diff = chain[i] - chain[i - 1];
    if (diff === 1) ones++;
    else if (diff === 3) threes++;
  }
  return ones * threes;
}
