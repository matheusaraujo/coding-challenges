use crate::{Answer, answer};

pub fn part1(puzzle_input: &[String]) -> Answer {
    let line = puzzle_input.first().map(|s| s.as_str()).unwrap_or("");

    let result = line.chars().fold(0i32, |acc, c| {
        match c {
            '(' => acc + 1,
            ')' => acc - 1,
            _ => acc,
        }
    });

    answer(result)
}
