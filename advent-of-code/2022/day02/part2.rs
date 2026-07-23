use crate::helpers::{Shape, outcome_score, round_columns};
use crate::{Any, any};

pub fn part2(puzzle_input: &[String]) -> Any {
    let total: u32 = puzzle_input
        .iter()
        .filter(|line| !line.is_empty())
        .map(|line| {
            let (opponent_c, outcome_c) = round_columns(line);
            let opponent = Shape::from_opponent(opponent_c);
            let you = match outcome_c {
                'X' => opponent.beats(),
                'Y' => opponent,
                'Z' => opponent.beaten_by(),
                _ => panic!("invalid outcome: {outcome_c}"),
            };
            you.score() + outcome_score(opponent, you)
        })
        .sum();
    any(total)
}
