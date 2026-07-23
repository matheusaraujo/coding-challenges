use crate::helpers::{max_geodes, parse_blueprints};
use crate::{Any, any};

pub fn part2(puzzle_input: &[String]) -> Any {
    let product: i64 = parse_blueprints(puzzle_input)
        .iter()
        .take(3)
        .map(|bp| max_geodes(bp, 32))
        .product();
    any(product)
}
