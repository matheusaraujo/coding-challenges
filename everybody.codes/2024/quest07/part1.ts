import { Device, parseInput } from "./helpers";

export function part1(puzzleInput: string[]): any {
  return simulate(parseInput(puzzleInput), 10);
}

function simulate(devices: Device[], segments: number): string {
  for (let i = 0; i < segments; i++) {
    for (const dev of devices) {
      const ac = dev.actions[i % dev.actions.length];
      if (ac === "+") dev.power++;
      else if (ac === "-") dev.power = Math.max(0, dev.power - 1);
      dev.acumulated += dev.power;
    }
  }
  return devices
    .sort((a, b) => b.acumulated - a.acumulated)
    .map((d) => d.id)
    .join("");
}
