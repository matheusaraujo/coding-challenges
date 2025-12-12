export type Device = {
  id: string;
  power: number;
  acumulated: number;
  actions: string[];
};

export function parseInput(puzzleInput: string[]): Device[] {
  return puzzleInput.map((line: string) => {
    const [id, actions] = line.split(":");
    return {
      id,
      power: 10,
      acumulated: 10,
      actions: actions.split(","),
    };
  });
}
