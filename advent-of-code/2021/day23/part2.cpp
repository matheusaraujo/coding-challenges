#include "helpers.h"
#include <any>
using namespace std;

// The folded part of the diagram adds two rows between the existing ones:
//   #D#C#B#A#
//   #D#B#A#C#
static const string EXTRA[4] = {"DD", "CB", "BA", "AC"};

any part2(const vector<string> &puzzleInput) {
  auto rooms = parseRooms(puzzleInput);
  for (int r = 0; r < 4; r++)
    rooms[r] = rooms[r][0] + EXTRA[r] + rooms[r][1];
  return minOrganizeEnergy(rooms);
}
