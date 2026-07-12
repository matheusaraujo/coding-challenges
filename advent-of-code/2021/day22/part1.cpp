#include "helpers.h"
#include <any>
using namespace std;

any part1(const vector<string> &puzzleInput) {
  vector<Step> initSteps;
  for (const auto &s : parseSteps(puzzleInput))
    if (s.x1 >= -50 && s.x2 <= 50 && s.y1 >= -50 && s.y2 <= 50 && s.z1 >= -50 &&
        s.z2 <= 50)
      initSteps.push_back(s);
  return countOnCubes(initSteps);
}
