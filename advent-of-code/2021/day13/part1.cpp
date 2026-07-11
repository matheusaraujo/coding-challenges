#include "helpers.h"
#include <any>
using namespace std;

any part1(const vector<string> &puzzleInput) {
  auto [dots, folds] = parseInput(puzzleInput);
  return (int)applyFold(dots, folds[0]).size();
}
