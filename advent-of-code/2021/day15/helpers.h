#pragma once
#include <string>
#include <vector>
using namespace std;

vector<vector<int>> parseGrid(const vector<string> &puzzleInput);

// Dijkstra from the top-left to the bottom-right corner.
int lowestTotalRisk(const vector<vector<int>> &grid);
