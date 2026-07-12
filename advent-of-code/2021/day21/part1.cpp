#include "helpers.h"
#include <any>
using namespace std;

any part1(const vector<string> &puzzleInput) {
  auto [p1, p2] = parseStartingPositions(puzzleInput);
  int pos[2] = {p1, p2}, score[2] = {0, 0};
  int die = 0, rolls = 0;

  for (int player = 0;; player = 1 - player) {
    int move = 0;
    for (int i = 0; i < 3; i++) {
      die = die % 100 + 1;
      move += die;
    }
    rolls += 3;

    pos[player] = (pos[player] + move - 1) % 10 + 1;
    score[player] += pos[player];
    if (score[player] >= 1000)
      return score[1 - player] * rolls;
  }
}
