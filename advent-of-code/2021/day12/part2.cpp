#include "helpers.h"
#include <any>
using namespace std;

any part2(const vector<string> &puzzleInput) {
  auto revisitOneSmallCave = [](const map<string, int> &visits,
                                const string &cave) {
    if (cave == "start")
      return false;
    for (const auto &[other, count] : visits)
      if (isSmall(other) && count >= 2)
        return false; // the single double-visit was already used
    return true;
  };
  return countPaths(puzzleInput, revisitOneSmallCave);
}
