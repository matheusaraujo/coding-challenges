from helpers import mod_inverse, parse_instructions


def part2(puzzle_input):
    n = 119315717514047
    iterations = 101741582076661
    pos = 2020

    # Get the single-pass transformation: f(x) = ax + b
    a, b = parse_instructions(puzzle_input, n)

    # We need to apply f(x) 'iterations' times.
    # Composing f(x) with itself k times gives:
    # A = a^k
    # B = b * (a^k - 1) / (a - 1)
    # Using geometric series formula: B = b * (A - 1) * inv(a - 1)

    final_a = pow(a, iterations, n)
    final_b = (b * (final_a - 1) * mod_inverse(a - 1, n)) % n

    # Part 2 asks: what card is at position 2020?
    # This means we need the inverse function: x = (pos - final_b) * inv(final_a)
    inv_a = mod_inverse(final_a, n)
    return ((pos - final_b) * inv_a) % n
