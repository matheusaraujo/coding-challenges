const MOD = 20201227;

export function part1(puzzleInput) {
  const [cardPub, doorPub] = puzzleInput.map(Number);
  let val = 1, loopSize = 0;
  while (val !== cardPub) {
    val = (val * 7) % MOD;
    loopSize++;
  }
  let key = 1;
  for (let i = 0; i < loopSize; i++) key = (key * doorPub) % MOD;
  return key;
}
