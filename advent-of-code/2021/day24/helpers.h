#pragma once
#include <string>
#include <tuple>
#include <vector>
using namespace std;

// MONAD treats z as a base-26 stack: blocks with "div z 1" push
// digit[i] + B, blocks with "div z 26" pop and require
// digit[j] = digit[i] + delta. Returns those (i, j, delta) constraints.
vector<tuple<int, int, int>>
digitConstraints(const vector<string> &puzzleInput);
