#pragma once
#include <string>
#include <vector>
using namespace std;

struct Step {
  bool on;
  long long x1, x2, y1, y2, z1, z2; // inclusive bounds
};

vector<Step> parseSteps(const vector<string> &puzzleInput);
long long countOnCubes(const vector<Step> &steps);
