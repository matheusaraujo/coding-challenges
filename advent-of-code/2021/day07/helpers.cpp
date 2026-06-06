#include "helpers.h"
#include <algorithm>
#include <climits>
#include <sstream>
using namespace std;

long long minFuel(const vector<string> &puzzleInput,
                  function<long long(long long)> cost) {
  vector<long long> positions;
  stringstream ss(puzzleInput[0]);
  string token;
  while (getline(ss, token, ','))
    positions.push_back(stoll(token));

  long long lo = *min_element(positions.begin(), positions.end());
  long long hi = *max_element(positions.begin(), positions.end());

  long long best = LLONG_MAX;
  for (long long target = lo; target <= hi; target++) {
    long long fuel = 0;
    for (long long p : positions)
      fuel += cost(abs(p - target));
    best = min(best, fuel);
  }
  return best;
}
