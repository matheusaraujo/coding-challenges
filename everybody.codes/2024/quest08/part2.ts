export function part2(puzzleInput: string[]): any {
  return solve(parseInt(puzzleInput[0]), 20240000);
}

function solve(p: number, n: number): number {
  const a = 1111;

  let totalBlocksUsed = 0;
  let k = 1;
  let prevThickness = 1;

  while (totalBlocksUsed <= n) {
    let currentThickness: number;
    if (k === 1) {
      currentThickness = 1;
    } else {
      currentThickness = (prevThickness * p) % a;
    }

    const currentWidth = 2 * k - 1;
    const blocksNeeded = currentWidth * currentThickness;

    if (totalBlocksUsed + blocksNeeded <= n) {
      totalBlocksUsed += blocksNeeded;
      prevThickness = currentThickness;
      k++;
    } else {
      break;
    }
  }

  const targetLayer = k;
  const targetWidth = 2 * targetLayer - 1;

  let targetThickness: number;
  if (targetLayer === 1) {
    targetThickness = 1;
  } else {
    if (k === 0) {
      targetThickness = 1;
    } else {
      targetThickness = (prevThickness * p) % a;
    }
  }

  const blocksNeededForTarget = targetWidth * targetThickness;

  const remainingAvailableBlocks = n - totalBlocksUsed;
  const missingBlocks = blocksNeededForTarget - remainingAvailableBlocks;

  return missingBlocks * targetWidth;
}
