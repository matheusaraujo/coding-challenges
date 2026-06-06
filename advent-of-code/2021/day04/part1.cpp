#include "helpers.h"
#include <any>
using namespace std;

any part1(const vector<string> &puzzleInput) {
  auto firstWinner = [](int) { return true; };
  return playBingo(puzzleInput, firstWinner);
}
