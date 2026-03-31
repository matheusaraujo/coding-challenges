import math
import re


def parse_reactions(puzzle_input):
    raw_data = (
        puzzle_input if isinstance(puzzle_input, str) else "\n".join(puzzle_input)
    )
    reactions = {}
    for line in raw_data.strip().split("\n"):
        # Matches all pairs of (amount, name)
        components = re.findall(r"(\d+) ([A-Z]+)", line)
        # The last component is the output, the rest are inputs
        *inputs, output = components
        out_qty, out_name = int(output[0]), output[1]
        reactions[out_name] = {
            "produce": out_qty,
            "ingredients": [(int(qty), name) for qty, name in inputs],
        }
    return reactions


def get_ore_needed(fuel_amount, reactions):
    needed = {"FUEL": fuel_amount}
    surplus = {}
    ore_count = 0

    while needed:
        chemical, amount = needed.popitem()
        if chemical == "ORE":
            ore_count += amount
            continue

        # Use surplus first
        from_surplus = min(amount, surplus.get(chemical, 0))
        amount -= from_surplus
        surplus[chemical] = surplus.get(chemical, 0) - from_surplus

        if amount <= 0:
            continue

        # Calculate how many times to run the reaction
        reaction = reactions[chemical]
        runs = math.ceil(amount / reaction["produce"])

        # Add ingredients to the 'needed' pile
        for ing_qty, ing_name in reaction["ingredients"]:
            needed[ing_name] = needed.get(ing_name, 0) + (ing_qty * runs)

        # Store extra produced
        extra = (runs * reaction["produce"]) - amount
        surplus[chemical] = surplus.get(chemical, 0) + extra

    return ore_count
