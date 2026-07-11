#pragma once
#include <set>
#include <string>
#include <utility>
#include <vector>
using namespace std;

using Dots = set<pair<int, int>>;
using Fold = pair<char, int>; // axis ('x' or 'y') and position

pair<Dots, vector<Fold>> parseInput(const vector<string> &puzzleInput);
Dots applyFold(const Dots &dots, const Fold &fold);
