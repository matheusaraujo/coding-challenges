#include "helpers.h"
#include <sstream>
using namespace std;

long long countFish(const vector<string> &puzzleInput, int days) {
  long long counts[9] = {};

  stringstream ss(puzzleInput[0]);
  string token;
  while (getline(ss, token, ','))
    counts[stoi(token)]++;

  for (int d = 0; d < days; d++) {
    long long spawning = counts[0];
    for (int i = 0; i < 8; i++)
      counts[i] = counts[i + 1];
    counts[8] = spawning;
    counts[6] += spawning;
  }

  long long total = 0;
  for (int i = 0; i <= 8; i++)
    total += counts[i];
  return total;
}
