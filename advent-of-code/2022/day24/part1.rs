use crate::helpers::{min_time, parse_blizzards};
use crate::{Any, any};

pub fn part1(puzzle_input: &[String]) -> Any {
    let bl = parse_blizzards(puzzle_input);
    any(min_time(&bl, bl.start, bl.end, 0))
}
