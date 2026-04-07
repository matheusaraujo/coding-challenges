from helpers import run_springscript


def part2(puzzle_input):
    # Logic: Jump if (Part 1 Logic) AND (Can move from D via walking E or jumping H)
    script = """NOT A J
NOT B T
OR T J
NOT C T
OR T J
AND D J
NOT E T
NOT T T
OR H T
AND T J
RUN
"""
    return run_springscript(puzzle_input, script)
