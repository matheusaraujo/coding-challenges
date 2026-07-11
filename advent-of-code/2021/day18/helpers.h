#pragma once
#include <string>
#include <vector>
using namespace std;

// A snailfish number as its leaves in order, each with its nesting depth.
struct Leaf {
  long long value;
  int depth;
};
using Number = vector<Leaf>;

Number parseNumber(const string &line);
Number add(const Number &a, const Number &b);
long long magnitude(const Number &number);
