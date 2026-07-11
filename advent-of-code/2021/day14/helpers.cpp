#include "helpers.h"
#include <algorithm>
#include <map>
using namespace std;

long long polymerScore(const vector<string> &puzzleInput, int steps) {
  const string &polymer = puzzleInput[0];

  map<pair<char, char>, char> rules;
  for (size_t i = 2; i < puzzleInput.size(); i++)
    rules[{puzzleInput[i][0], puzzleInput[i][1]}] = puzzleInput[i][6];

  map<pair<char, char>, long long> pairs;
  for (size_t i = 0; i + 1 < polymer.size(); i++)
    pairs[{polymer[i], polymer[i + 1]}]++;

  for (int s = 0; s < steps; s++) {
    map<pair<char, char>, long long> next;
    for (const auto &[pair, count] : pairs) {
      auto [a, b] = pair;
      char inserted = rules.at(pair);
      next[{a, inserted}] += count;
      next[{inserted, b}] += count;
    }
    pairs = next;
  }

  // Each element is the first char of exactly one pair, except the last
  // element of the polymer, which never changes.
  map<char, long long> counts;
  for (const auto &[pair, count] : pairs)
    counts[pair.first] += count;
  counts[polymer.back()]++;

  long long mostCommon = 0, leastCommon = 0;
  for (const auto &[_, count] : counts) {
    mostCommon = max(mostCommon, count);
    leastCommon = leastCommon == 0 ? count : min(leastCommon, count);
  }
  return mostCommon - leastCommon;
}
