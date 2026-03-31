from helpers import get_ore_needed, parse_reactions


def part1(puzzle_input):
    reactions = parse_reactions(puzzle_input)
    return get_ore_needed(1, reactions)
