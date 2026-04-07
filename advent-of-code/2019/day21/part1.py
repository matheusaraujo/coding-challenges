from helpers import run_springscript


def part1(puzzle_input):
    # Logic: Jump if (NOT A OR NOT B OR NOT C) AND D
    script = """NOT A J
NOT B T
OR T J
NOT C T
OR T J
AND D J
WALK
"""
    return run_springscript(puzzle_input, script)
