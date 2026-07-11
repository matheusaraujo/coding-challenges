#include "helpers.h"
#include <any>
using namespace std;

any part1(const vector<string> &puzzleInput) {
  auto neverRevisit = [](const map<string, int> &, const string &) {
    return false;
  };
  return countPaths(puzzleInput, neverRevisit);
}
