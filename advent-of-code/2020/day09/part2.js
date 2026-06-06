import { findInvalid } from "./helpers.js";

export function part2(puzzleInput) {
  const nums = puzzleInput.map(Number);
  const target = findInvalid(nums, 25);
  for (let i = 0; i < nums.length; i++) {
    let sum = 0;
    for (let j = i; j < nums.length; j++) {
      sum += nums[j];
      if (sum === target && j > i) {
        const range = nums.slice(i, j + 1);
        return Math.min(...range) + Math.max(...range);
      }
      if (sum > target) break;
    }
  }
}
