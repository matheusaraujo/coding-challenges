export function part1(puzzleInput) {
  let password = "abcdefgh".split("");

  for (const line of puzzleInput) {
    let match;

    if ((match = line.match(/swap position (\d+) with position (\d+)/))) {
      const x = Number(match[1]);
      const y = Number(match[2]);
      [password[x], password[y]] = [password[y], password[x]];
      continue;
    }

    if ((match = line.match(/swap letter (\w) with letter (\w)/))) {
      const x = match[1];
      const y = match[2];
      password = password.map((c) => (c === x ? y : c === y ? x : c));
      continue;
    }

    if ((match = line.match(/rotate (left|right) (\d+) step/))) {
      const dir = match[1];
      const steps = Number(match[2]) % password.length;
      if (dir === "left") {
        password = password.slice(steps).concat(password.slice(0, steps));
      } else {
        password = password.slice(-steps).concat(password.slice(0, -steps));
      }
      continue;
    }

    if ((match = line.match(/rotate based on position of letter (\w)/))) {
      const x = match[1];
      const idx = password.indexOf(x);
      let steps = 1 + idx;
      if (idx >= 4) steps++;
      steps %= password.length;
      password = password.slice(-steps).concat(password.slice(0, -steps));
      continue;
    }

    if ((match = line.match(/reverse positions (\d+) through (\d+)/))) {
      const x = Number(match[1]);
      const y = Number(match[2]);
      const reversed = password.slice(x, y + 1).reverse();
      password.splice(x, y - x + 1, ...reversed);
      continue;
    }

    if ((match = line.match(/move position (\d+) to position (\d+)/))) {
      const x = Number(match[1]);
      const y = Number(match[2]);
      const [char] = password.splice(x, 1);
      password.splice(y, 0, char);
      continue;
    }
  }

  return password.join("");
}
