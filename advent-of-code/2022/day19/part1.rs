use crate::helpers::{max_geodes, parse_blueprints};
use crate::{Any, any};

pub fn part1(puzzle_input: &[String]) -> Any {
    let total: i64 = parse_blueprints(puzzle_input)
        .iter()
        .map(|bp| bp.id * max_geodes(bp, 24))
        .sum();
    any(total)
}
