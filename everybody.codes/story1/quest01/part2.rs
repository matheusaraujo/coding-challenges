use super::helpers::eni;
use super::helpers::solve;

pub fn part2(puzzle_input: &[String]) -> String {
    solve(puzzle_input, |n, e, m| {
        let exp = e.saturating_sub(5);
        let base_score = mod_pow(n, exp, m);
        eni(base_score, n, 5, m)
    })
    .to_string()
}

fn mod_pow(mut base: usize, mut exp: usize, modulus: usize) -> usize {
    if modulus == 1 {
        return 0;
    }
    let mut res = 1;
    base %= modulus;
    while exp > 0 {
        if exp % 2 == 1 {
            res = (res * base) % modulus;
        }
        exp >>= 1;
        base = (base * base) % modulus;
    }
    res
}
