import { Device, parseInput } from "./helpers";

export function part2(puzzleInput: string[]): string {
  const racetrack =
    "-=++=-==++=++=-=+=-=+=+=--=-=++=-==++=-+=-=+=-=+=+=++=-+==++=++=-=-=---=++==--+++==++=+=--==++==+++=++=+++=--=+=-=+=-+=-+=-+-=+=-=+=-+++=+==++++==---=+=+=-S";
  return simulate(parseInput(puzzleInput), racetrack, 10);
}

function simulate(devices: Device[], racetrack: string, loops: number): string {
  for (let i = 0; i < loops; i++) {
    for (let j = 0; j < racetrack.length; j++) {
      const rt = racetrack[j];
      if (rt === "S") continue;
      for (const dev of devices) {
        const ac = dev.actions[(i * racetrack.length + j) % dev.actions.length];
        if (rt === "+") dev.power++;
        else if (rt === "-") dev.power = Math.max(0, dev.power - 1);
        else {
          if (ac === "+") dev.power++;
          else if (ac === "-") dev.power = Math.max(0, dev.power - 1);
        }
        dev.acumulated += dev.power;
      }
    }
  }
  return devices
    .sort((a, b) => b.acumulated - a.acumulated)
    .map((d) => d.id)
    .join("");
}
