#include "helpers.h"
#include <cstdio>
using namespace std;

vector<int> peakHeightsOfAllHits(const vector<string> &puzzleInput) {
  int xMin, xMax, yMin, yMax;
  sscanf(puzzleInput[0].c_str(), "target area: x=%d..%d, y=%d..%d", &xMin,
         &xMax, &yMin, &yMax);

  vector<int> peaks;
  for (int vx0 = 1; vx0 <= xMax; vx0++)
    for (int vy0 = yMin; vy0 <= -yMin; vy0++) {
      int x = 0, y = 0, vx = vx0, vy = vy0, peak = 0;
      while (x <= xMax && y >= yMin) {
        if (x >= xMin && y <= yMax) {
          peaks.push_back(peak);
          break;
        }
        x += vx;
        y += vy;
        peak = max(peak, y);
        if (vx > 0)
          vx--;
        vy--;
      }
    }

  return peaks;
}
