pub fn parse_numbers(puzzle_input: &[String]) -> Vec<i64> {
    puzzle_input
        .iter()
        .filter(|line| !line.is_empty())
        .map(|line| line.parse().unwrap())
        .collect()
}

pub fn mix(numbers: &[i64], rounds: usize) -> Vec<i64> {
    let len = numbers.len();
    let mut next: Vec<usize> = (0..len).map(|i| (i + 1) % len).collect();
    let mut prev: Vec<usize> = (0..len).map(|i| (i + len - 1) % len).collect();

    for _ in 0..rounds {
        for i in 0..len {
            let shift = numbers[i].rem_euclid((len - 1) as i64);
            if shift == 0 {
                continue;
            }

            let p = prev[i];
            let n = next[i];
            next[p] = n;
            prev[n] = p;

            let mut target = p;
            for _ in 0..shift {
                target = next[target];
            }
            let after = next[target];
            next[target] = i;
            prev[i] = target;
            next[i] = after;
            prev[after] = i;
        }
    }

    let start = numbers
        .iter()
        .position(|&v| v == 0)
        .expect("value 0 exists");
    let mut result = Vec::with_capacity(len);
    let mut cur = start;
    for _ in 0..len {
        result.push(numbers[cur]);
        cur = next[cur];
    }
    result
}

pub fn grove_coordinates(mixed: &[i64]) -> i64 {
    let len = mixed.len();
    [1000, 2000, 3000]
        .iter()
        .map(|&offset| mixed[offset % len])
        .sum()
}
