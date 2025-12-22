use crate::{Any, any};

pub fn part1(puzzle_input: &[String]) -> Any {
    let line = puzzle_input.first().map(|s| s.as_str()).unwrap_or("");

    let result:i32 = line.chars().fold(0i32, |acc, c| {
        match c {
            '(' => acc + 1,
            ')' => acc - 1,
            _ => acc,
        }
    });

    any(result)
}
