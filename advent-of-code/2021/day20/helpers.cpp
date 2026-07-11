#include "helpers.h"
using namespace std;

long long litPixelsAfter(const vector<string> &puzzleInput, int steps) {
  const string &algorithm = puzzleInput[0];
  vector<string> image(puzzleInput.begin() + 2, puzzleInput.end());

  // The infinite background: all dark initially, but it flips every step
  // when algorithm[0] is '#' (and algorithm[511] is '.').
  char background = '.';

  for (int s = 0; s < steps; s++) {
    int rows = image.size(), cols = image[0].size();
    vector<string> next(rows + 2, string(cols + 2, '.'));

    for (int r = 0; r < rows + 2; r++)
      for (int c = 0; c < cols + 2; c++) {
        int index = 0;
        for (int dr = -1; dr <= 1; dr++)
          for (int dc = -1; dc <= 1; dc++) {
            int nr = r - 1 + dr, nc = c - 1 + dc;
            char pixel = nr >= 0 && nr < rows && nc >= 0 && nc < cols
                             ? image[nr][nc]
                             : background;
            index = index * 2 + (pixel == '#');
          }
        next[r][c] = algorithm[index];
      }

    image = next;
    background = algorithm[background == '#' ? 511 : 0];
  }

  long long lit = 0;
  for (const auto &row : image)
    for (char pixel : row)
      lit += pixel == '#';
  return lit;
}
