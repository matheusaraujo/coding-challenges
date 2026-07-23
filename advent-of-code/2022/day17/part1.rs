use crate::helpers::{parse_jets, simulate};
use crate::{Any, any};

pub fn part1(puzzle_input: &[String]) -> Any {
    let jets = parse_jets(puzzle_input);
    any(simulate(&jets, 2022))
}
