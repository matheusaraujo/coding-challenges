import { bestAsteroid, parseAsteroids } from "./helpers.js";

export function part2(puzzleInput) {
  const { bestAsteroid: station } = bestAsteroid(puzzleInput);
  const [sy, sx] = station;

  const asteroids = parseAsteroids(puzzleInput);
  const asteroidList = Object.keys(asteroids)
    .map((k) => k.split(",").map(Number))
    .filter(([y, x]) => !(y === sy && x === sx));

  const angleMap = new Map();

  for (const [y, x] of asteroidList) {
    const dy = y - sy;
    const dx = x - sx;

    let angle = Math.atan2(dx, -dy);
    if (angle < 0) angle += 2 * Math.PI;

    const distance = Math.sqrt(dy * dy + dx * dx);

    if (!angleMap.has(angle)) angleMap.set(angle, []);
    angleMap.get(angle).push({ y, x, distance });
  }

  for (const arr of angleMap.values()) {
    arr.sort((a, b) => a.distance - b.distance);
  }

  const sortedAngles = Array.from(angleMap.keys()).sort((a, b) => a - b);

  const vaporized = [];
  while (vaporized.length < asteroidList.length) {
    for (const angle of sortedAngles) {
      const group = angleMap.get(angle);
      if (group.length > 0) vaporized.push(group.shift());
    }
  }

  const asteroid200 = vaporized[199];
  return asteroid200.x * 100 + asteroid200.y;
}
