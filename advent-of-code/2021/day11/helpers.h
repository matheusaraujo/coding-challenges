#pragma once
#include <string>
#include <vector>
using namespace std;

vector<vector<int>> parseGrid(const vector<string> &puzzleInput);

// Advances the grid one step and returns how many octopuses flashed.
int step(vector<vector<int>> &grid);
