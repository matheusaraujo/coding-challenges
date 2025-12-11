function solve(puzzleInput, func) {
  const arr = puzzleInput.map((e) => +e).sort((a, b) => a - b);
  const target = arr[func(arr)];
  let result = 0;
  for (let i = 0; i < arr.length; i++) {
    result += Math.abs(target - arr[i]);
  }
  return result;
}

module.exports = solve;
