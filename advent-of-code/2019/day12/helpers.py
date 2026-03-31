import math
import re


def get_lcm(a, b):
    return abs(a * b) // math.gcd(a, b)


def get_lcm_list(numbers):
    res = numbers[0]
    for i in range(1, len(numbers)):
        res = get_lcm(res, numbers[i])
    return res


def parse_moons(puzzle_input):
    # Ensure input is a string
    raw_data = (
        puzzle_input if isinstance(puzzle_input, str) else "\n".join(puzzle_input)
    )

    moons = []
    for line in raw_data.strip().split("\n"):
        if not line.strip():
            continue
        # Find all sequences of digits, potentially starting with a minus sign
        # This matches the x, y, and z values directly
        nums = re.findall(r"-?\d+", line)
        if nums:
            moons.append({"pos": [int(n) for n in nums], "vel": [0, 0, 0]})
    return moons
