#include "pb_helpers.h"

void *part2(char **puzzle_input) {
  int floor = 0;
  for (int i = 0; puzzle_input[0][i] != '\0'; i++) {
    floor += puzzle_input[0][i] == '(' ? 1 : -1;
    if (floor == -1) {
      return answer(i + 1);
    }
  }
  return answer(0);
}