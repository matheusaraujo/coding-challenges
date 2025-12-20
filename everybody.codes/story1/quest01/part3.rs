use super::helpers::solve;
use std::collections::HashMap;

pub fn part3(puzzle_input: &[String]) -> String {
    solve(puzzle_input, find_cycle).to_string()
}

fn find_cycle(n: usize, e: usize, m: usize) -> usize {
    let mut score = 1;
    let mut total = 0;
    let mut index = 0;

    let mut sums = Vec::with_capacity(m);
    let mut seen = HashMap::with_capacity(m);

    sums.push(0);

    loop {
        score = (score * n) % m;
        total += score;
        index += 1;

        if let Some(previous) = seen.insert(score, index) {
            let cycle_length = index - previous;
            let cycle_total = total - sums[previous];

            let remaining = e - index + 1;
            let quotient = remaining / cycle_length;
            let remainder = remaining % cycle_length;

            return (total - score)
                + (quotient * cycle_total)
                + (sums[previous + remainder - 1] - sums[previous - 1]);
        }

        sums.push(total);
    }
}
