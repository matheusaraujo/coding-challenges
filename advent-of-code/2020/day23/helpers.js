export function play(cups, moves) {
  const n = cups.length;
  const next = new Int32Array(n + 1);
  for (let i = 0; i < n - 1; i++) next[cups[i]] = cups[i + 1];
  next[cups[n - 1]] = cups[0];

  let cur = cups[0];
  for (let m = 0; m < moves; m++) {
    const a = next[cur], b = next[a], c = next[b];
    next[cur] = next[c];
    let dest = cur === 1 ? n : cur - 1;
    while (dest === a || dest === b || dest === c) {
      dest = dest === 1 ? n : dest - 1;
    }
    next[c] = next[dest];
    next[dest] = a;
    cur = next[cur];
  }
  return next;
}
