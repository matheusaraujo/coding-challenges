export function tokenize(line) {
  return line.replace(/\s/g, "").split(
    /(?<=\d)(?=\D)|(?<=\D)(?=\d)|(?<=\D)(?=\D)/g,
  )
    .flatMap((t) => (isNaN(t) ? [t] : [Number(t)]));
}

// Part 1: left-to-right, all ops equal precedence
export function evalP1(tokens, pos) {
  let val = atom1(tokens, pos);
  while (pos[0] < tokens.length && tokens[pos[0]] !== ")") {
    const op = tokens[pos[0]++];
    const right = atom1(tokens, pos);
    val = op === "+" ? val + right : val * right;
  }
  return val;
}

function atom1(tokens, pos) {
  if (tokens[pos[0]] === "(") {
    pos[0]++;
    const val = evalP1(tokens, pos);
    pos[0]++; // skip ")"
    return val;
  }
  return tokens[pos[0]++];
}

// Part 2: + has higher precedence than *
export function evalP2(tokens, pos) {
  let val = addExpr(tokens, pos);
  while (pos[0] < tokens.length && tokens[pos[0]] === "*") {
    pos[0]++;
    val *= addExpr(tokens, pos);
  }
  return val;
}

function addExpr(tokens, pos) {
  let val = atom2(tokens, pos);
  while (pos[0] < tokens.length && tokens[pos[0]] === "+") {
    pos[0]++;
    val += atom2(tokens, pos);
  }
  return val;
}

function atom2(tokens, pos) {
  if (tokens[pos[0]] === "(") {
    pos[0]++;
    const val = evalP2(tokens, pos);
    pos[0]++; // skip ")"
    return val;
  }
  return tokens[pos[0]++];
}
