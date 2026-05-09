export function parsePassports(puzzleInput) {
  const passports = [];
  let current = {};
  for (const line of puzzleInput) {
    if (line === "") {
      passports.push(current);
      current = {};
      continue;
    }
    for (const pair of line.split(" ")) {
      const [key, value] = pair.split(":");
      current[key] = value;
    }
  }
  if (Object.keys(current).length) passports.push(current);
  return passports;
}

export const REQUIRED = ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"];

export function hasRequiredFields(passport) {
  return REQUIRED.every((key) => key in passport);
}
