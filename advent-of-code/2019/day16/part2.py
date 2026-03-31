def part2(puzzle_input):
    digits = [int(c) for c in puzzle_input[0].strip()]
    offset = int(puzzle_input[0][:7])

    signal = (digits * 10000)[offset:]

    for _ in range(100):
        suffix_sum = 0
        for i in range(len(signal) - 1, -1, -1):
            suffix_sum = (suffix_sum + signal[i]) % 10
            signal[i] = suffix_sum

    return "".join(map(str, signal[:8]))
