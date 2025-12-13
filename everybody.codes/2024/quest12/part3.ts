export function part3(puzzleInput: string[]): any {
  const meteors = puzzleInput.map((line) => ({
    x: parseInt(line.split(" ")[0]),
    y: parseInt(line.split(" ")[1]),
  }));

  const maxX = Math.max(...meteors.map((m) => m.x));
  const traces: number[][][] = new Array(3).fill("").map(() => []);
  for (let i = 0; i < 3; i++) {
    traces[i].push([]);
    for (let p = 1; p < Math.floor(maxX / 2); ++p) {
      const t = [i];
      const current = { x: 0, y: i };

      for (let j = 0; j < p; j++) {
        current.x++;
        current.y++;
        t.push(current.y);
      }

      for (let j = 0; j < p; j++) {
        current.x++;
        t.push(current.y);
      }

      while (current.y > 0) {
        current.x++;
        current.y--;
        t.push(current.y);
      }

      traces[i].push(t);
    }
  }

  return meteors.reduce((sum, meteor) => {
    let time = Math.round(meteor.x / 2);
    meteor.x -= time;
    meteor.y -= time;

    let min = Infinity;
    while (meteor.x >= 0 && meteor.y >= 0) {
      for (let i = 0; i < 3; i++) {
        for (let p = 1; p < traces[i].length; p++) {
          const score = (i + 1) * p;
          if (score > min) break;

          const powerList = traces[i][p];
          if (meteor.x < powerList.length && powerList[meteor.x] === meteor.y) {
            min = Math.min(min, score);
            return sum + min;
          }
        }
      }
      time++;
      meteor.x--;
      meteor.y--;
    }

    return sum;
  }, 0);
}
