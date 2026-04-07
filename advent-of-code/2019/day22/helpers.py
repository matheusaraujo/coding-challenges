def mod_inverse(a, n):
    """Computes the modular multiplicative inverse using Fermat's Little Theorem."""
    return pow(a, n - 2, n)


def parse_instructions(puzzle_input, n):
    """
    Translates all instructions into a single linear function: f(x) = ax + b % n.
    Returns (a, b).
    """
    if isinstance(puzzle_input, str):
        lines = puzzle_input.strip().split("\n")
    else:
        lines = [line.strip() for line in puzzle_input if line.strip()]

    a, b = 1, 0
    for line in lines:
        if "deal into new stack" in line:
            # New index = -old_index - 1
            l_a, l_b = -1, -1
        elif "cut" in line:
            # New index = old_index - k
            k = int(line.split()[-1])
            l_a, l_b = 1, -k
        elif "deal with increment" in line:
            # New index = old_index * k
            k = int(line.split()[-1])
            l_a, l_b = k, 0
        else:
            continue

        # Compose: f(g(x)) where g(x) = ax + b and f(x) = l_a*x + l_b
        # f(g(x)) = l_a*(ax + b) + l_b = (l_a*a)x + (l_a*b + l_b)
        a = (l_a * a) % n
        b = (l_a * b + l_b) % n

    return a, b
