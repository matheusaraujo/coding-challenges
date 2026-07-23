use crate::helpers::{Shape, outcome_score, round_columns};
use crate::{Any, any};

pub fn part1(puzzle_input: &[String]) -> Any {
    let total: u32 = puzzle_input
        .iter()
        .filter(|line| !line.is_empty())
        .map(|line| {
            let (opponent_c, you_c) = round_columns(line);
            let opponent = Shape::from_opponent(opponent_c);
            let you = match you_c {
                'X' => Shape::Rock,
                'Y' => Shape::Paper,
                'Z' => Shape::Scissors,
                _ => panic!("invalid your shape: {you_c}"),
            };
            you.score() + outcome_score(opponent, you)
        })
        .sum();
    any(total)
}
