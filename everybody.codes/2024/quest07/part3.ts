const SIGN: Record<string, number> = {
  "+": 1,
  "-": -1,
  "=": 0,
  S: 0,
};

const DIRS = [
  [0, -1],
  [1, 0],
  [0, 1],
  [-1, 0],
];

const planValid = (plan: string[]) =>
  plan.filter((v) => v == "+").length == 5 &&
  plan.filter((v) => v == "-").length == 3 &&
  plan.filter((v) => v == "=").length == 3;

const race = (
  powers: number[],
  vals: string[],
  segments: number,
  racetrack: string[],
) => {
  for (let seg = 0; seg < segments; seg++) {
    let pow = 10;
    if (seg > 0) pow = powers[seg - 1];
    let action = SIGN[racetrack[seg % racetrack.length]];
    if (action == 0) action = SIGN[vals[seg % vals.length]];
    powers[seg] = Math.max(0, pow + action);
  }
  return powers.reduce((a, v) => a + v, 0);
};

const run3 = (vals: string[], racetrack: string[], loops: number) =>
  race([], vals, racetrack.length * loops, racetrack);

const track3 = `S+= +=-== +=++=     =+=+=--=    =-= ++=     +=-  =+=++=-+==+ =++=-=-=--
- + +   + =   =     =      =   == = - -     - =  =         =-=        -
= + + +-- =-= ==-==-= --++ +  == == = +     - =  =    ==++=    =++=-=++
+ + + =     +         =  + + == == ++ =     = =  ==   =   = =++=
= = + + +== +==     =++ == =+=  =  +  +==-=++ =   =++ --= + =
+ ==- = + =   = =+= =   =       ++--          +     =   = = =--= ==++==
=     ==- ==+-- = = = ++= +=--      ==+ ==--= +--+=-= ==- ==   =+=    =
-               = = = =   +  +  ==+ = = +   =        ++    =          -
-               = + + =   +  -  = + = = +   =        +     =          -
--==++++==+=+++-= =-= =-+-=  =+-= =-= =--   +=++=+++==     -=+=++==+++-`;

const buildTrack = (trackLit: string) => {
  const map = trackLit.split("\n").map((line) => line.split(""));
  let [prevX, prevY] = [0, 0],
    [curX, curY] = [1, 0],
    res = "";

  while (map[curY][curX] != "S") {
    res += map[curY][curX];
    DIRS.some(([dx, dy]) => {
      const [x, y] = [curX + dx, curY + dy];
      if (x == prevX && y == prevY) return false;
      if (y < 0 || y > map.length - 1 || x < 0 || x > map[0].length - 1)
        return false;
      if (map[y][x] == " " || map[y][x] == undefined) return false;
      [prevX, prevY] = [curX, curY];
      [curX, curY] = [x, y];
      return true;
    });
  }
  return (res + "S").split("");
};

export function part3(puzzleInput: string[0]): any {
  const NUM2SIGN = ["+", "-", "="];

  const laps = 11;

  const track = buildTrack(track3),
    toBeat = run3(puzzleInput[0].split(":")[1].split(","), track, laps);
  let res = 0;

  for (let i = 0; i < Math.pow(3, 11); i++) {
    const plan = i
      .toString(3)
      .padStart(11, "0")
      .split("")
      .map((v) => NUM2SIGN[v]);
    if (planValid(plan) && run3(plan, track, laps) > toBeat) res++;
  }
  return res;
}
