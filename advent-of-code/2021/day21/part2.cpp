#include "helpers.h"
#include <algorithm>
#include <any>
#include <map>
#include <tuple>
using namespace std;

// How many of the 27 three-roll outcomes sum to 3..9.
static const int ROLL_COUNTS[10] = {0, 0, 0, 1, 3, 6, 7, 6, 3, 1};

using State = tuple<int, int, int, int>; // current player's pos/score, then
                                         // the other player's
using Wins = pair<long long, long long>; // wins for current player / other

static map<State, Wins> cache;

static Wins countWins(int pos, int score, int otherPos, int otherScore) {
  State state{pos, score, otherPos, otherScore};
  auto it = cache.find(state);
  if (it != cache.end())
    return it->second;

  Wins wins{0, 0};
  for (int roll = 3; roll <= 9; roll++) {
    int newPos = (pos + roll - 1) % 10 + 1;
    int newScore = score + newPos;
    if (newScore >= 21) {
      wins.first += ROLL_COUNTS[roll];
    } else {
      auto [otherWins, myWins] =
          countWins(otherPos, otherScore, newPos, newScore);
      wins.first += ROLL_COUNTS[roll] * myWins;
      wins.second += ROLL_COUNTS[roll] * otherWins;
    }
  }

  cache[state] = wins;
  return wins;
}

any part2(const vector<string> &puzzleInput) {
  auto [p1, p2] = parseStartingPositions(puzzleInput);
  auto [p1Wins, p2Wins] = countWins(p1, 0, p2, 0);
  return max(p1Wins, p2Wins);
}
