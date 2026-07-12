#include "helpers.h"
#include <map>
#include <queue>
using namespace std;

vector<string> parseRooms(const vector<string> &puzzleInput) {
  vector<string> rooms(4);
  for (int room = 0; room < 4; room++) {
    rooms[room] += puzzleInput[2][3 + 2 * room];
    rooms[room] += puzzleInput[3][3 + 2 * room];
  }
  return rooms;
}

// Dijkstra over burrow states: 11 hallway cells then the rooms top to
// bottom. Legal moves are room -> hallway spot and hallway -> home room.
long long minOrganizeEnergy(const vector<string> &rooms) {
  const int depth = rooms[0].size();
  const long long COST[4] = {1, 10, 100, 1000};
  const int DOOR[4] = {2, 4, 6, 8};
  const int SPOTS[7] = {0, 1, 3, 5, 7, 9, 10};

  string start(11, '.'), goal(11, '.');
  for (int r = 0; r < 4; r++) {
    start += rooms[r];
    goal += string(depth, 'A' + r);
  }

  auto pathClear = [](const string &state, int from, int to) {
    for (int x = min(from, to); x <= max(from, to); x++)
      if (x != from && state[x] != '.')
        return false;
    return true;
  };

  priority_queue<pair<long long, string>, vector<pair<long long, string>>,
                 greater<>>
      pq;
  map<string, long long> dist;
  pq.push({0, start});
  dist[start] = 0;

  while (!pq.empty()) {
    auto [energy, state] = pq.top();
    pq.pop();
    if (state == goal)
      return energy;
    if (energy > dist[state])
      continue;

    auto tryMove = [&](string next, long long cost) {
      auto it = dist.find(next);
      if (it == dist.end() || energy + cost < it->second) {
        dist[next] = energy + cost;
        pq.push({energy + cost, next});
      }
    };

    // Topmost amphipod of each unsettled room -> any reachable hallway spot
    for (int r = 0; r < 4; r++) {
      int base = 11 + r * depth, top = 0;
      while (top < depth && state[base + top] == '.')
        top++;
      if (top == depth)
        continue;
      bool settled = true;
      for (int s = top; s < depth; s++)
        settled &= state[base + s] == 'A' + r;
      if (settled)
        continue;

      char c = state[base + top];
      for (int spot : SPOTS) {
        if (state[spot] != '.' || !pathClear(state, DOOR[r], spot))
          continue;
        string next = state;
        next[spot] = c;
        next[base + top] = '.';
        tryMove(next, (top + 1 + abs(spot - DOOR[r])) * COST[c - 'A']);
      }
    }

    // Hallway amphipod -> its room, if the room has no strangers
    for (int spot : SPOTS) {
      char c = state[spot];
      if (c == '.')
        continue;
      int r = c - 'A', base = 11 + r * depth, bottom = -1;
      bool strangers = false;
      for (int s = 0; s < depth; s++) {
        if (state[base + s] == '.')
          bottom = s;
        else
          strangers |= state[base + s] != c;
      }
      if (strangers || !pathClear(state, spot, DOOR[r]))
        continue;

      string next = state;
      next[spot] = '.';
      next[base + bottom] = c;
      tryMove(next, (abs(spot - DOOR[r]) + bottom + 1) * COST[r]);
    }
  }

  return -1;
}
