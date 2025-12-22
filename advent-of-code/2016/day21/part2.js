export function part2(puzzleInput) {
  let password = "fbgdceah".split("");

  for (let i = puzzleInput.length - 1; i >= 0; i--) {
    const line = puzzleInput[i];
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

    if ((match = line.match(/rotate left (\d+) step/))) {
      const steps = Number(match[1]) % password.length;
      password = password.slice(-steps).concat(password.slice(0, -steps));
      continue;
    }

    if ((match = line.match(/rotate right (\d+) step/))) {
      const steps = Number(match[1]) % password.length;
      password = password.slice(steps).concat(password.slice(0, steps));
      continue;
    }

    if ((match = line.match(/rotate based on position of letter (\w)/))) {
      const x = match[1];
      const len = password.length;
      let original;

      for (let j = 0; j < len; j++) {
        const test = password.slice(j).concat(password.slice(0, j));
        const idx = test.indexOf(x);
        let steps = 1 + idx + (idx >= 4 ? 1 : 0);
        steps %= len;
        const rotated = test.slice(-steps).concat(test.slice(0, -steps));
        if (rotated.join("") === password.join("")) {
          original = test;
          break;
        }
      }
      password = original;
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
      const [char] = password.splice(y, 1);
      password.splice(x, 0, char);
      continue;
    }
  }

  return password.join("");
}
