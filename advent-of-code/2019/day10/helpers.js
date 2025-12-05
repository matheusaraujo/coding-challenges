function parseAsteroids(puzzleInput) {
  const a = {};
  for (let y = 0; y < puzzleInput.length; y++) {
    for (let x = 0; x < puzzleInput[y].length; x++) {
      if (puzzleInput[y][x] === "#") a[`${y},${x}`] = true;
    }
  }
  return a;
}

function gcd(a, b) {
  a = Math.abs(a);
  b = Math.abs(b);
  while (b !== 0) [a, b] = [b, a % b];
  return a;
}

function bestAsteroid(puzzleInput) {
  const asteroids = parseAsteroids(puzzleInput);
  const coords = Object.keys(asteroids).map((k) => k.split(",").map(Number));

  let maxVisible = 0;
  let best = null;

  for (let i = 0; i < coords.length; i++) {
    const [y1, x1] = coords[i];
    const directions = new Set();

    for (let j = 0; j < coords.length; j++) {
      if (i === j) continue;

      const [y2, x2] = coords[j];
      let dy = y2 - y1;
      let dx = x2 - x1;

      const g = gcd(dy, dx);
      dy /= g;
      dx /= g;

      directions.add(`${dy},${dx}`);
    }

    if (directions.size > maxVisible) {
      maxVisible = directions.size;
      best = [y1, x1];
    }
  }

  return { bestAsteroid: best, maxVisible };
}

module.exports = { parseAsteroids, gcd, bestAsteroid };
