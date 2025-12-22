use super::helpers::parse_input;
use crate::{Any, any};
use std::collections::HashMap;

pub fn part2(puzzle_input: &[String]) -> Any {
    let (left, right) = parse_input(puzzle_input);

    let mut counts = HashMap::new();
    for &num in &right {
        *counts.entry(num).or_insert(0) += 1;
    }

    let result: i32 = left.iter()
        .map(|&l| l * counts.get(&l).unwrap_or(&0))
        .sum();

    any(result)
}
