#include <any>
#include <string>
#include <vector>
using namespace std;

static int filter(vector<string> nums, bool mostCommon) {
  int bits = nums[0].size();
  for (int b = 0; b < bits && nums.size() > 1; b++) {
    int ones = 0;
    for (const auto &s : nums)
      if (s[b] == '1')
        ones++;
    int zeros = nums.size() - ones;
    char keep =
        mostCommon ? (ones >= zeros ? '1' : '0') : (zeros <= ones ? '0' : '1');
    vector<string> next;
    for (const auto &s : nums)
      if (s[b] == keep)
        next.push_back(s);
    nums = next;
  }
  return stoi(nums[0], nullptr, 2);
}

any part2(const vector<string> &puzzleInput) {
  return filter(puzzleInput, true) * filter(puzzleInput, false);
}
