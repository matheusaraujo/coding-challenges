#include "helpers.h"
using namespace std;

pair<int, int> parseStartingPositions(const vector<string> &puzzleInput) {
  return {stoi(puzzleInput[0].substr(puzzleInput[0].find(": ") + 2)),
          stoi(puzzleInput[1].substr(puzzleInput[1].find(": ") + 2))};
}
