#include "helpers.h"
#include <algorithm>
#include <any>
#include <map>
using namespace std;

any part2(const vector<string> &puzzleInput) {
  const map<char, long long> points{{'(', 1}, {'[', 2}, {'{', 3}, {'<', 4}};

  vector<long long> scores;
  string stack;
  for (const auto &line : puzzleInput) {
    if (firstIllegalChar(line, stack))
      continue; // corrupted lines are discarded

    long long score = 0;
    for (auto it = stack.rbegin(); it != stack.rend(); ++it)
      score = score * 5 + points.at(*it);
    scores.push_back(score);
  }

  sort(scores.begin(), scores.end());
  return scores[scores.size() / 2];
}
