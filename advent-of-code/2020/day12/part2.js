export function part2(puzzleInput) {
  let sx = 0, sy = 0, wx = 10, wy = 1; // waypoint relative to ship
  for (const line of puzzleInput) {
    const action = line[0], val = Number(line.slice(1));
    if (action === "N") wy += val;
    else if (action === "S") wy -= val;
    else if (action === "E") wx += val;
    else if (action === "W") wx -= val;
    else if (action === "F") {
      sx += wx * val;
      sy += wy * val;
    } else {
      const times = ((action === "L" ? 360 - val : val) / 90) % 4;
      for (let i = 0; i < times; i++) [wx, wy] = [wy, -wx];
    }
  }
  return Math.abs(sx) + Math.abs(sy);
}
