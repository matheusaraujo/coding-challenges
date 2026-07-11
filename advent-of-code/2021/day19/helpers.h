#pragma once
#include <array>
#include <set>
#include <string>
#include <vector>
using namespace std;

using Point = array<int, 3>;

// Everything expressed in scanner 0's coordinate frame.
struct Alignment {
  set<Point> beacons;
  vector<Point> scanners;
};

Alignment alignScanners(const vector<string> &puzzleInput);
