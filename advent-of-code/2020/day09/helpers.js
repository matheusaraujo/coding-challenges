export function findInvalid(nums, preamble) {
  for (let i = preamble; i < nums.length; i++) {
    const window = nums.slice(i - preamble, i);
    const target = nums[i];
    const valid = window.some((a, ai) =>
      window.some((b, bi) => ai !== bi && a + b === target)
    );
    if (!valid) return target;
  }
}
