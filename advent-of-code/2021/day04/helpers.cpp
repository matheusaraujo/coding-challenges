#include "helpers.h"
#include <sstream>
using namespace std;

struct Board {
  int nums[5][5];
  bool marked[5][5];

  void mark(int n) {
    for (int r = 0; r < 5; r++)
      for (int c = 0; c < 5; c++)
        if (nums[r][c] == n)
          marked[r][c] = true;
  }

  bool wins() const {
    for (int i = 0; i < 5; i++) {
      bool row = true, col = true;
      for (int j = 0; j < 5; j++) {
        if (!marked[i][j])
          row = false;
        if (!marked[j][i])
          col = false;
      }
      if (row || col)
        return true;
    }
    return false;
  }

  int score(int last) const {
    int sum = 0;
    for (int r = 0; r < 5; r++)
      for (int c = 0; c < 5; c++)
        if (!marked[r][c])
          sum += nums[r][c];
    return sum * last;
  }
};

int playBingo(const vector<string> &puzzleInput, bool findLast) {
  vector<int> draws;
  stringstream ss(puzzleInput[0]);
  string token;
  while (getline(ss, token, ','))
    draws.push_back(stoi(token));

  vector<Board> boards;
  for (size_t i = 2; i < puzzleInput.size(); i += 6) {
    Board b = {};
    for (int r = 0; r < 5; r++) {
      istringstream row(puzzleInput[i + r]);
      for (int c = 0; c < 5; c++)
        row >> b.nums[r][c];
    }
    boards.push_back(b);
  }

  vector<bool> won(boards.size(), false);
  int remaining = boards.size();
  for (int d : draws)
    for (size_t i = 0; i < boards.size(); i++) {
      if (won[i])
        continue;
      boards[i].mark(d);
      if (boards[i].wins()) {
        won[i] = true;
        if (!findLast || --remaining == 0)
          return boards[i].score(d);
      }
    }
  return -1;
}
