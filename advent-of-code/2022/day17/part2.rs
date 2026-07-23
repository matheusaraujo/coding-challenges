use crate::helpers::{parse_jets, simulate};
use crate::{Any, any};

pub fn part2(puzzle_input: &[String]) -> Any {
    let jets = parse_jets(puzzle_input);
    any(simulate(&jets, 1_000_000_000_000))
}
