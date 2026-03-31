# helpers.py

BASE_PATTERN = [0, 1, 0, -1]


def build_pattern(position, length):
    pattern = []
    for p in BASE_PATTERN:
        pattern.extend([p] * position)
    # repeat pattern until long enough
    while len(pattern) <= length:
        pattern *= 2
    # skip the very first value
    return pattern[1 : length + 1]


def apply_phase(signal):
    output = []
    length = len(signal)

    for i in range(length):
        pattern = build_pattern(i + 1, length)
        value = sum(a * b for a, b in zip(signal, pattern))
        output.append(abs(value) % 10)

    return output
