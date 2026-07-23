use crate::helpers::{best_scores, parse_problem};
use crate::{Any, any};

pub fn part2(puzzle_input: &[String]) -> Any {
    let problem = parse_problem(puzzle_input);
    let best = best_scores(&problem, 26);
    let entries: Vec<(u32, i64)> = best.into_iter().collect();

    let max_total = entries
        .iter()
        .flat_map(|&(mask1, score1)| {
            entries
                .iter()
                .filter_map(move |&(mask2, score2)| (mask1 & mask2 == 0).then_some(score1 + score2))
        })
        .max()
        .unwrap_or(0);
    any(max_total)
}
