#include "helpers.h"
#include <algorithm>
using namespace std;

Number parseNumber(const string &line) {
  Number number;
  int depth = 0;
  for (size_t i = 0; i < line.size(); i++) {
    if (line[i] == '[')
      depth++;
    else if (line[i] == ']')
      depth--;
    else if (isdigit(line[i]))
      number.push_back({line[i] - '0', depth});
  }
  return number;
}

// The leftmost pair nested inside four pairs (leaves at depth 5) explodes.
static bool explode(Number &n) {
  for (size_t i = 0; i + 1 < n.size(); i++) {
    if (n[i].depth == 5) {
      if (i > 0)
        n[i - 1].value += n[i].value;
      if (i + 2 < n.size())
        n[i + 2].value += n[i + 1].value;
      n[i] = {0, 4};
      n.erase(n.begin() + i + 1);
      return true;
    }
  }
  return false;
}

// The leftmost value >= 10 splits into a pair one level deeper.
static bool split(Number &n) {
  for (size_t i = 0; i < n.size(); i++) {
    if (n[i].value >= 10) {
      long long value = n[i].value;
      int depth = n[i].depth + 1;
      n[i] = {value / 2, depth};
      n.insert(n.begin() + i + 1, {(value + 1) / 2, depth});
      return true;
    }
  }
  return false;
}

Number add(const Number &a, const Number &b) {
  Number sum;
  for (const auto &leaf : a)
    sum.push_back({leaf.value, leaf.depth + 1});
  for (const auto &leaf : b)
    sum.push_back({leaf.value, leaf.depth + 1});

  while (explode(sum) || split(sum))
    ;
  return sum;
}

long long magnitude(const Number &number) {
  Number n = number;
  while (n.size() > 1) {
    int deepest = 0;
    for (const auto &leaf : n)
      deepest = max(deepest, leaf.depth);
    for (size_t i = 0; i + 1 < n.size(); i++) {
      if (n[i].depth == deepest && n[i + 1].depth == deepest) {
        n[i] = {3 * n[i].value + 2 * n[i + 1].value, deepest - 1};
        n.erase(n.begin() + i + 1);
        break;
      }
    }
  }
  return n[0].value;
}
