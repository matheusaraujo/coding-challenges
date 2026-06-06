#include "helpers.h"
#include <algorithm>
#include <map>
#include <sstream>
using namespace std;

static bool isSubset(const string &a, const string &b) {
  for (char c : a)
    if (b.find(c) == string::npos)
      return false;
  return true;
}

static string sorted(string s) {
  sort(s.begin(), s.end());
  return s;
}

static pair<vector<string>, vector<string>> parseLine(const string &line) {
  auto sep = line.find(" | ");
  stringstream ps(line.substr(0, sep));
  stringstream os(line.substr(sep + 3));
  vector<string> patterns, outputs;
  string token;
  while (ps >> token)
    patterns.push_back(token);
  while (os >> token)
    outputs.push_back(token);
  return {patterns, outputs};
}

int countEasyDigits(const vector<string> &puzzleInput) {
  int count = 0;
  for (const auto &line : puzzleInput) {
    auto [_, outputs] = parseLine(line);
    for (const auto &o : outputs) {
      int len = o.size();
      if (len == 2 || len == 3 || len == 4 || len == 7)
        count++;
    }
  }
  return count;
}

long long sumOutputValues(const vector<string> &puzzleInput) {
  long long total = 0;

  for (const auto &line : puzzleInput) {
    auto [patterns, outputs] = parseLine(line);

    // Normalize everything
    for (auto &p : patterns)
      p = sorted(p);

    for (auto &o : outputs)
      o = sorted(o);

    string d0, d1, d2, d3, d4, d5, d6, d7, d8, d9;
    vector<string> len5;
    vector<string> len6;

    for (const auto &p : patterns) {
      switch (p.size()) {
      case 2:
        d1 = p;
        break;
      case 3:
        d7 = p;
        break;
      case 4:
        d4 = p;
        break;
      case 7:
        d8 = p;
        break;
      case 5:
        len5.push_back(p);
        break;
      case 6:
        len6.push_back(p);
        break;
      }
    }

    // 0,6,9
    for (const auto &p : len6) {
      if (!isSubset(d1, p))
        d6 = p;
      else if (isSubset(d4, p))
        d9 = p;
      else
        d0 = p;
    }

    // 2,3,5
    for (const auto &p : len5) {
      if (isSubset(d1, p))
        d3 = p;
      else if (isSubset(p, d6))
        d5 = p;
      else
        d2 = p;
    }

    map<string, int> lookup{{d0, 0}, {d1, 1}, {d2, 2}, {d3, 3}, {d4, 4},
                            {d5, 5}, {d6, 6}, {d7, 7}, {d8, 8}, {d9, 9}};

    int value = 0;
    for (const auto &o : outputs)
      value = value * 10 + lookup.at(o); // throws if deduction failed

    total += value;
  }

  return total;
}
