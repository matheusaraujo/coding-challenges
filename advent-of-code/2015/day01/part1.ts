const countOccurrences: (str: string, char: string) => number = (str, char) => {
  return str.split(char).length - 1;
};

export function part1(puzzleInput: string[]): any {
  return (
    countOccurrences(puzzleInput[0], "(") -
    countOccurrences(puzzleInput[0], ")")
  );
}
