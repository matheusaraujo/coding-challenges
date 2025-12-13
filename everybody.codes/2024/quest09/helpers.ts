export function dp(targetBrightness: number, stamps: number[]): number[] {
  const dp = new Array(targetBrightness + 1).fill(Infinity);
  dp[0] = 0;
  for (let i = 1; i <= targetBrightness; i++) {
    for (const stamp of stamps) {
      if (i - stamp >= 0) {
        dp[i] = Math.min(dp[i], dp[i - stamp] + 1);
      }
    }
  }
  return dp;
}
