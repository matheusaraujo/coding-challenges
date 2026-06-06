export function parse(puzzleInput) {
  const sep = puzzleInput.indexOf("");
  return [
    puzzleInput.slice(1, sep).map(Number),
    puzzleInput.slice(sep + 2).map(Number),
  ];
}

export function score(deck) {
  return deck.reduce((sum, card, i) => sum + card * (deck.length - i), 0);
}

export function playRecursive(d1, d2) {
  const seen = new Set();
  while (d1.length && d2.length) {
    const key = `${d1}|${d2}`;
    if (seen.has(key)) return [1, d1];
    seen.add(key);
    const [a, b] = [d1.shift(), d2.shift()];
    let winner;
    if (d1.length >= a && d2.length >= b) {
      [winner] = playRecursive(d1.slice(0, a), d2.slice(0, b));
    } else {
      winner = a > b ? 1 : 2;
    }
    winner === 1 ? d1.push(a, b) : d2.push(b, a);
  }
  return d1.length ? [1, d1] : [2, d2];
}
