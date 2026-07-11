#include "helpers.h"
#include <cstdio>
#include <map>
using namespace std;

using Rotation = array<array<int, 3>, 3>;

// The 24 orientation-preserving axis rotations: signed axis permutations
// with determinant +1.
static vector<Rotation> makeRotations() {
  const int perms[6][3] = {{0, 1, 2}, {0, 2, 1}, {1, 0, 2},
                           {1, 2, 0}, {2, 0, 1}, {2, 1, 0}};
  const int parity[6] = {1, -1, -1, 1, 1, -1};

  vector<Rotation> rotations;
  for (int p = 0; p < 6; p++)
    for (int sx : {1, -1})
      for (int sy : {1, -1})
        for (int sz : {1, -1}) {
          if (parity[p] * sx * sy * sz != 1)
            continue;
          const int signs[3] = {sx, sy, sz};
          Rotation m{};
          for (int i = 0; i < 3; i++)
            m[i][perms[p][i]] = signs[i];
          rotations.push_back(m);
        }
  return rotations;
}

static Point rotate(const Point &p, const Rotation &m) {
  Point r;
  for (int i = 0; i < 3; i++)
    r[i] = m[i][0] * p[0] + m[i][1] * p[1] + m[i][2] * p[2];
  return r;
}

// Tries to overlap `raw` (in its own frame) with `known` (already in scanner
// 0's frame): >= 12 beacons must coincide under one rotation + translation.
static bool match(const vector<Point> &known, const vector<Point> &raw,
                  const vector<Rotation> &rotations, vector<Point> &alignedOut,
                  Point &scannerOut) {
  for (const auto &rot : rotations) {
    vector<Point> rotated;
    for (const auto &p : raw)
      rotated.push_back(rotate(p, rot));

    map<Point, int> offsets;
    for (const auto &k : known)
      for (const auto &b : rotated)
        offsets[{k[0] - b[0], k[1] - b[1], k[2] - b[2]}]++;

    for (const auto &[offset, count] : offsets) {
      if (count < 12)
        continue;
      alignedOut.clear();
      for (const auto &b : rotated)
        alignedOut.push_back(
            {b[0] + offset[0], b[1] + offset[1], b[2] + offset[2]});
      scannerOut = offset;
      return true;
    }
  }
  return false;
}

Alignment alignScanners(const vector<string> &puzzleInput) {
  vector<vector<Point>> scans;
  for (const auto &line : puzzleInput) {
    if (line.rfind("--- scanner", 0) == 0)
      scans.push_back({});
    else if (!line.empty()) {
      Point p;
      sscanf(line.c_str(), "%d,%d,%d", &p[0], &p[1], &p[2]);
      scans.back().push_back(p);
    }
  }

  const auto rotations = makeRotations();
  int n = scans.size();
  vector<vector<Point>> aligned(n);
  vector<bool> done(n, false);

  Alignment result;
  result.scanners.assign(n, {0, 0, 0});
  aligned[0] = scans[0];
  done[0] = true;

  vector<int> queue{0};
  while (!queue.empty()) {
    int ref = queue.back();
    queue.pop_back();
    for (int i = 0; i < n; i++)
      if (!done[i] && match(aligned[ref], scans[i], rotations, aligned[i],
                            result.scanners[i])) {
        done[i] = true;
        queue.push_back(i);
      }
  }

  for (const auto &beacons : aligned)
    result.beacons.insert(beacons.begin(), beacons.end());
  return result;
}
