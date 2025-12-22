use super::helpers::parse_input;
use crate::{Answer, answer};
use std::collections::HashMap;

pub fn part2(puzzle_input: &[String]) -> Answer {
    let (left, right) = parse_input(puzzle_input);

    let mut counts = HashMap::new();
    for &num in &right {
        *counts.entry(num).or_insert(0) += 1;
    }

    let result: i32 = left.iter()
        .map(|&l| l * counts.get(&l).unwrap_or(&0))
        .sum();

    answer(result)
}
