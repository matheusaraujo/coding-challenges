use std::collections::HashMap;

struct Parameters {
    a: usize,
    b: usize,
    c: usize,
    x: usize,
    y: usize,
    z: usize,
    m: usize,
}

pub fn solve(puzzle_input: &[String], wrapper: fn(usize, usize, usize) -> usize) -> usize {
    puzzle_input
        .iter()
        .map(|s| parse_row(s))
        .map(|r| wrapper(r.a, r.x, r.m) + wrapper(r.b, r.y, r.m) + wrapper(r.c, r.z, r.m))
        .max()
        .unwrap()
}

fn parse_row(line: &str) -> Parameters {
    let v: HashMap<&str, usize> = line
        .split_whitespace()
        .map(|kv| {
            let (k, v) = kv.split_once('=').unwrap();
            (k, v.parse().unwrap())
        })
        .collect();

    Parameters {
        a: v["A"],
        b: v["B"],
        c: v["C"],
        x: v["X"],
        y: v["Y"],
        z: v["Z"],
        m: v["M"],
    }
}

pub fn eni(mut score: usize, n: usize, e: usize, m: usize) -> usize {
    let mut result = 0;
    let mut power = 0;

    for _ in 0..e {
        score = (score * n) % m;
        result += score * 10_usize.pow(power);
        power += if score < 10 { 1 } else { score.ilog10() + 1 };
    }

    result
}
