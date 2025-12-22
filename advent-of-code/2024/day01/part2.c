enum { MAX_SIZE = 100000 };

#include "helpers.h"
#include "pb_helpers.h"

void *part2(char **puzzle_input, int size) {
  int result = 0;
  int left[size], right[size];
  int count[MAX_SIZE];

  pb_memset(count, 0, sizeof(count));

  parse_and_sort(puzzle_input, size, left, right);

  for (int i = 0; i < size; i++) {
    count[right[i]]++;
  }

  for (int i = 0; i < size; i++) {
    result += left[i] * count[left[i]];
  }

  return any(result);
}
