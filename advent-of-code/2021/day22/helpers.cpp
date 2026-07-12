#include "helpers.h"
#include <algorithm>
#include <cstdio>
using namespace std;

vector<Step> parseSteps(const vector<string> &puzzleInput) {
  vector<Step> steps;
  for (const auto &line : puzzleInput) {
    Step s;
    s.on = line[1] == 'n';
    sscanf(line.c_str() + (s.on ? 3 : 4),
           "x=%lld..%lld,y=%lld..%lld,z=%lld..%lld", &s.x1, &s.x2, &s.y1, &s.y2,
           &s.z1, &s.z2);
    steps.push_back(s);
  }
  return steps;
}

// Inclusion-exclusion over signed cuboids: every new step cancels its
// overlap with each cuboid placed so far, then "on" steps add themselves.
long long countOnCubes(const vector<Step> &steps) {
  struct Cuboid {
    long long x1, x2, y1, y2, z1, z2, sign;
  };
  vector<Cuboid> placed;

  for (const auto &s : steps) {
    vector<Cuboid> added;
    for (const auto &c : placed) {
      Cuboid i{max(s.x1, c.x1), min(s.x2, c.x2), max(s.y1, c.y1),
               min(s.y2, c.y2), max(s.z1, c.z1), min(s.z2, c.z2),
               -c.sign};
      if (i.x1 <= i.x2 && i.y1 <= i.y2 && i.z1 <= i.z2)
        added.push_back(i);
    }
    if (s.on)
      added.push_back({s.x1, s.x2, s.y1, s.y2, s.z1, s.z2, 1});
    placed.insert(placed.end(), added.begin(), added.end());
  }

  long long total = 0;
  for (const auto &c : placed)
    total += c.sign * (c.x2 - c.x1 + 1) * (c.y2 - c.y1 + 1) * (c.z2 - c.z1 + 1);
  return total;
}
