export function part2(puzzleInput) {
  const map = { x: 0, A: 0, B: 1, C: 3, D: 5 };
  const data = puzzleInput[0].split("");
  let result = data.map((x) => map[x]).reduce((x, y) => x + y);
  for (let i = 0; i < data.length; i += 2) {
    if (data[i] !== "x" && data[i + 1] !== "x") {
      result += 2;
    }
  }
  return result;
}
