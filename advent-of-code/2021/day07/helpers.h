#pragma once
#include <functional>
#include <string>
#include <vector>
using namespace std;

long long minFuel(const vector<string> &puzzleInput,
                  function<long long(long long)> cost);
