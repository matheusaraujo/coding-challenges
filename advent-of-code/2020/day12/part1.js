export function part1(puzzleInput) {
  let x = 0, y = 0, dir = 0; // dir: 0=E,90=S,180=W,270=N
  const dx = [1, 0, -1, 0], dy = [0, -1, 0, 1];
  for (const line of puzzleInput) {
    const action = line[0], val = Number(line.slice(1));
    if (action === "N") y += val;
    else if (action === "S") y -= val;
    else if (action === "E") x += val;
    else if (action === "W") x -= val;
    else if (action === "L") dir = (dir + 360 - val) % 360;
    else if (action === "R") dir = (dir + val) % 360;
    else if (action === "F") {
      x += dx[dir / 90] * val;
      y += dy[dir / 90] * val;
    }
  }
  return Math.abs(x) + Math.abs(y);
}
