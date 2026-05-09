import { hasRequiredFields, parsePassports } from "./helpers.js";

const VALIDATORS = {
  byr: (v) => /^\d{4}$/.test(v) && +v >= 1920 && +v <= 2002,
  iyr: (v) => /^\d{4}$/.test(v) && +v >= 2010 && +v <= 2020,
  eyr: (v) => /^\d{4}$/.test(v) && +v >= 2020 && +v <= 2030,
  hgt: (v) => {
    const m = v.match(/^(\d+)(cm|in)$/);
    if (!m) return false;
    const n = +m[1];
    return m[2] === "cm" ? n >= 150 && n <= 193 : n >= 59 && n <= 76;
  },
  hcl: (v) => /^#[0-9a-f]{6}$/.test(v),
  ecl: (v) => ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"].includes(v),
  pid: (v) => /^\d{9}$/.test(v),
};

export function part2(puzzleInput) {
  return parsePassports(puzzleInput)
    .filter(hasRequiredFields)
    .filter((p) => Object.entries(VALIDATORS).every(([k, fn]) => fn(p[k])))
    .length;
}
