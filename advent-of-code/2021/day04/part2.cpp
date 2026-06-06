#include "helpers.h"
#include <any>
using namespace std;

any part2(const vector<string> &puzzleInput) {
  auto lastWinner = [](int remainingBoards) { return remainingBoards == 0; };
  return playBingo(puzzleInput, lastWinner);
}
