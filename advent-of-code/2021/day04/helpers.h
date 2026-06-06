#pragma once
#include <functional>
#include <string>
#include <vector>
using namespace std;

// Decides, right after a board wins, whether its score is the answer,
// given how many boards have not won yet.
using WinPolicy = function<bool(int remainingBoards)>;

int playBingo(const vector<string> &puzzleInput, const WinPolicy &isTarget);
