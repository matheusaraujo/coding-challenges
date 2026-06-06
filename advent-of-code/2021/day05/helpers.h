#pragma once
#include <functional>
#include <string>
#include <vector>
using namespace std;

int countOverlaps(const vector<string> &puzzleInput,
                  function<bool(int, int, int, int)> include);
