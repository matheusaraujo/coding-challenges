#include "helpers.h"
using namespace std;

pair<Dots, vector<Fold>> parseInput(const vector<string> &puzzleInput) {
  Dots dots;
  vector<Fold> folds;

  for (const auto &line : puzzleInput) {
    if (line.empty())
      continue;
    if (line.rfind("fold along ", 0) == 0) {
      char axis = line[11];
      folds.push_back({axis, stoi(line.substr(13))});
    } else {
      auto sep = line.find(',');
      dots.insert({stoi(line.substr(0, sep)), stoi(line.substr(sep + 1))});
    }
  }

  return {dots, folds};
}

Dots applyFold(const Dots &dots, const Fold &fold) {
  auto [axis, pos] = fold;
  Dots folded;
  for (auto [x, y] : dots) {
    if (axis == 'x' && x > pos)
      x = 2 * pos - x;
    else if (axis == 'y' && y > pos)
      y = 2 * pos - y;
    folded.insert({x, y});
  }
  return folded;
}
