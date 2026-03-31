from helpers import parse_reactions
from part1 import get_ore_needed


def part2(puzzle_input):
    reactions = parse_reactions(puzzle_input)
    trillion = 1000000000000

    # Low: at least 1 fuel. High: a massive number.
    low = 1
    high = trillion  # Safe upper bound
    ans = 0

    while low <= high:
        mid = (low + high) // 2
        if get_ore_needed(mid, reactions) <= trillion:
            ans = mid
            low = mid + 1
        else:
            high = mid - 1

    return ans
