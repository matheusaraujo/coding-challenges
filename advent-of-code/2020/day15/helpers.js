export function memoryGame(puzzleInput, target) {
  const nums = puzzleInput[0].split(",").map(Number);
  const last = new Map();
  nums.slice(0, -1).forEach((n, i) => last.set(n, i + 1));
  let prev = nums.at(-1);
  for (let turn = nums.length + 1; turn <= target; turn++) {
    const next = last.has(prev) ? turn - 1 - last.get(prev) : 0;
    last.set(prev, turn - 1);
    prev = next;
  }
  return prev;
}
